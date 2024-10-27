package grpc

import (
	"service/internal/service"
	desc "service/pkg/grpc/aggregator_v1"
)

type Server struct {
	desc.UnimplementedAggregatorServer
	Service *service.Service
}

func NewServer(Service *service.Service) *Server {
	return &Server{
		Service: Service,
	}
}
