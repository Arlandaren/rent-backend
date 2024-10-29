package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"os"
	"os/signal"
	"service/internal/repository"
	"service/internal/service"
	"service/internal/shared/kafka"
	"service/internal/shared/storage/postgres"
	"sync"
	"syscall"

	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	transport "service/internal/transport/grpc"

	desc "service/pkg/grpc/aggregator_v1"
)

const grpcAddress = ":8080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}

	postgresCfg, err := postgres.GetConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	kafkaBrokers := []string{os.Getenv("KAFKA_BROKER_ADDRESS")}
	if kafkaBrokers[0] == "" {
		kafkaBrokers = []string{"localhost:9092"}
	}
	kafkaCfg := kafka.NewConfig(kafkaBrokers)

	kafkaProducer, err := kafka.NewProducer(kafkaCfg)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	defer kafkaProducer.Close()

	dbPool, err := postgres.InitPostgres(postgresCfg, 5)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	defer dbPool.Close()

	pgxWrapper := postgres.NewWrapper(dbPool)

	appRepository := repository.NewRepository(pgxWrapper)

	appService := service.NewService(appRepository, kafkaProducer)

	appServer := transport.NewServer(appService)

	aggregator := service.NewAggregator(appService, nil)
	consumer := kafka.NewConsumer(aggregator)
	aggregator.SetConsumer(consumer)

	wg.Add(1)
	go func() {
		defer wg.Done()
		group := "aggregator_group"
		topics := []string{
			"booking_created",
			"booking_begin",
			"booking_updated",
			"booking_finished",
			"booking_cancelled",
			"apartment_created",
			"apartment_removed",
			"apartment_updated",
			"customer_created",
			"customer_removed",
			"customer_updated",
		}

		if err := aggregator.Start(ctx, kafkaBrokers, group, topics); err != nil {
			log.Fatalf("Failed to start Kafka consumer: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startGrpcServer(ctx, appServer); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down servers...")

	cancel()
	wg.Wait()
	log.Println("Servers gracefully stopped.")
}

func startGrpcServer(ctx context.Context, appServer *transport.Server) error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	desc.RegisterAggregatorServer(grpcServer, appServer)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	<-ctx.Done()

	grpcServer.GracefulStop()
	return nil
}
