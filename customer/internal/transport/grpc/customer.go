package grpc

import (
	"service/internal/service"
	desc "service/pkg/grpc/customer_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	desc.CustomerServiceServer
	service *service.Service
}

func NewServer(customerService *service.Service) *Server {
	return &Server{
		service: customerService,
	}
}

func (s *Server) NewCustomer(ctx context.Context, req *desc.NewCustomerRequest) (*desc.NewCustomerResponse, error) {
	customer, err := s.service.New(ctx, req.GetName(), req.GetPhone(), req.GetPassport())

	if err != nil {
		log.Printf("Failed to create new customer: %v\n", err)

		return nil, status.Errorf(codes.Internal, "failed to create customer: %v", err)
	}

	newCustomer := &desc.Customer{
		Id:        int64(customer.ID),
		Name:      customer.Name,
		Phone:     customer.Phone,
		Passport:  customer.Passport,
		CreatedAt: timestamppb.New(customer.CreatedAt),
	}

	log.Printf("Created new customer: %+v\n", customer)

	return &desc.NewCustomerResponse{
		Customer: newCustomer,
	}, nil
}

func (s *Server) RemoveCustomer(ctx context.Context, req *desc.RemoveCustomerRequest) (*desc.RemoveCustomerResponse, error) {
	err := s.service.Remove(ctx, int(req.GetId()))
	if err != nil {
		log.Printf("Failed to remove customer: %v\n", err)
		return nil, status.Errorf(codes.Internal, "failed to remove customer: %v", err)
	}
	return &desc.RemoveCustomerResponse{Success: true}, nil
}

func (s *Server) UpdateCustomer(ctx context.Context, req *desc.UpdateCustomerRequest) (*desc.UpdateCustomerResponse, error) {
	err := s.service.Update(ctx, int(req.GetId()), req.GetName(), req.GetPhone(), req.GetPassport())
	if err != nil {
		log.Printf("Failed to update customer: %v\n", err)
		return nil, status.Errorf(codes.Internal, "failed to update customer: %v", err)
	}

	return &desc.UpdateCustomerResponse{Success: true}, nil

}
