package main

import (
	"context"
	"customer_service/internal/config"
	"customer_service/internal/service"
	"customer_service/internal/shared/storage/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc/credentials/insecure"

	transport "customer_service/internal/transport/grpc"

	desc "customer_service/pkg/grpc/customer_v1"
)

const grpcAddress = ":8080"

func main() {
	cfg, err := config.GetPostgres()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	dbPool, err := postgres.InitPostgres(cfg, 5)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer dbPool.Close()

	customerService := service.NewCustomerService(dbPool)

	customerServer := transport.NewCustomerServer(customerService)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := startGrpcServer(ctx, customerServer); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
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

func startGrpcServer(ctx context.Context, customerServer *transport.CustomerServer) error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	// Предполагается, что у вас есть реализация сервера, которая принимает dbPool
	desc.RegisterCustomerServiceServer(grpcServer, customerServer)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at %v\n", grpcAddress)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	<-ctx.Done()

	grpcServer.GracefulStop()
	return nil
}