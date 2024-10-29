package grpc

import (
	"context"
	"service/internal/service"
	desc "service/pkg/grpc/aggregator_v1"
)

type Server struct {
	desc.AggregatorServer
	Service *service.Service
}

func NewServer(Service *service.Service) *Server {
	return &Server{
		Service: Service,
	}
}

func (s *Server) CustomerGet(ctx context.Context, request *desc.CustomerGetRequest) (*desc.CustomerGetResponse, error) {
	return s.Service.CustomerGet(ctx, request)
}

func (s *Server) CustomerList(ctx context.Context, request *desc.CustomerListRequest) (*desc.CustomerListResponse, error) {
	return s.Service.CustomerList(ctx, request)
}

func (s *Server) ApartmentGet(ctx context.Context, request *desc.ApartmentGetRequest) (*desc.ApartmentGetResponse, error) {
	return s.Service.ApartmentGet(ctx, request)
}

func (s *Server) ApartmentList(ctx context.Context, request *desc.ApartmentListRequest) (*desc.ApartmentListResponse, error) {
	return s.Service.ApartmentList(ctx, request)
}

func (s *Server) BookingGet(ctx context.Context, request *desc.BookGetRequest) (*desc.BookGetResponse, error) {
	return s.Service.BookingGet(ctx, request)
}

func (s *Server) BookingList(ctx context.Context, request *desc.BookListRequest) (*desc.BookListResponse, error) {
	return s.Service.BookingList(ctx, request)
}
