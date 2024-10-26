package grpc

import (
	"service/internal/service"
	desc "service/pkg/grpc/booking_v1"
)

type Server struct {
	desc.UnimplementedBookingServiceServer
	Service *service.Service
}

func NewServer(Service *service.Service) *Server {
	return &Server{
		Service: Service,
	}
}
