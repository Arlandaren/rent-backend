package grpc

import (
	"context"
	"customer_service/internal/service"
	desc "customer_service/pkg/grpc/customer_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type CustomerServer struct {
	desc.UnimplementedCustomerServiceServer
	customerService *service.CustomerService
}

func NewCustomerServer(customerService *service.CustomerService) *CustomerServer {
	return &CustomerServer{
		customerService: customerService,
	}
}

func (s *CustomerServer) NewCustomer(ctx context.Context, req *desc.NewCustomerRequest) (*desc.NewCustomerResponse, error) {

	customer, err := s.customerService.NewCustomer(req.GetName(), req.GetPhone(), req.GetPassport())
	if err != nil {
		return nil, err
	}
	newCustomer := &desc.Customer{
		Name:      customer,
		Phone:     customer,
		Passport:  customer,
		CreatedAt: timestamppb.Now(),
	}
	log.Printf("Created new customer: %+v\n", customer)

	return &desc.NewCustomerResponse{
		Customer: newCustomer,
	}, nil
}
