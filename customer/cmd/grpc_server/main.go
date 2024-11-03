package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"service/internal/repository"
	"service/internal/service"
	"service/internal/shared/kafka"
	"service/internal/shared/storage/postgres"
	"time"

	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc/credentials/insecure"

	transport "service/internal/transport/grpc"

	desc "service/pkg/grpc/customer_v1"
)

const (
	grpcAddress = ":50051"
	httpAddress = ":8086"
)

func main() {
	postgresCfg, err := postgres.GetConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	kafkaCfg := kafka.NewConfig([]string{os.Getenv("KAFKA_BROKER_ADDRESS")})

	kafkaProducer, err := kafka.NewProducer(kafkaCfg)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	dbPool, err := postgres.InitPostgres(postgresCfg, 5)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	defer dbPool.Close()

	pgxWrapper := postgres.NewWrapper(dbPool)

	appRepository := repository.NewRepository(pgxWrapper)

	appService := service.NewService(appRepository, kafkaProducer)

	appServer := transport.NewServer(appService)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := startGrpcServer(ctx, appServer); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := startHttpServer(ctx); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gRPC server...")

	cancel()
	wg.Wait()
	log.Println("Server gracefully stopped.")
}

func startGrpcServer(ctx context.Context, appServer *transport.Server) error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	desc.RegisterCustomerServiceServer(grpcServer, appServer)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	go func() {
		if err := grpcServer.Serve(list); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Printf("gRPC server failed: %v", err)
		}
	}()

	log.Printf("gRPC server listening at %v\n", grpcAddress)

	<-ctx.Done()

	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	return nil
}

func startHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterCustomerServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("failed to register service handler: %w", err)
	}

	handler := AllowCORS(mux)

	srv := &http.Server{
		Addr:    httpAddress,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server exited with error: %v", err)
		}
	}()

	log.Printf("HTTP server listening at %v\n", httpAddress)

	<-ctx.Done()

	log.Println("Shutting down HTTP server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("HTTP server Shutdown failed: %w", err)
	}

	return nil
}
