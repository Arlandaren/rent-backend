package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"service/internal/service"
	desc "service/pkg/grpc/apartment_v1"
)

type Server struct {
	desc.UnimplementedApartmentServiceServer
	Service *service.Service
}

func NewServer(Service *service.Service) *Server {
	return &Server{
		Service: Service,
	}
}

func (s *Server) NewApartment(ctx context.Context, req *desc.NewApartmentRequest) (*desc.NewApartmentResponse, error) {

	apartment, err := s.Service.New(ctx, req.GetTitle(), req.GetExpenses())
	if err != nil {
		log.Printf("Failed to create new customer: %v\n", err)

		return nil, status.Errorf(codes.Internal, "failed to create customer: %v", err)
	}

	newApartment := &desc.Apartment{
		Id:        apartment.ID,
		Title:     apartment.Title,
		Expenses:  apartment.Expenses,
		CreatedAt: timestamppb.New(apartment.CreatedAt),
	}

	return &desc.NewApartmentResponse{
		Apartment: newApartment,
	}, nil
}

func (s *Server) RemoveApartment(ctx context.Context, req *desc.RemoveApartmentRequest) (*desc.RemoveApartmentResponse, error) {
	err := s.Service.Remove(ctx, req.GetId())
	if err != nil {
		log.Printf("Failed to remove customer: %v\n", err)

		return nil, status.Errorf(codes.Internal, "failed to remove customer: %v", err)
	}

	return &desc.RemoveApartmentResponse{
		Success: true,
	}, nil
}
