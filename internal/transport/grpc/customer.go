package grpc

import (
	"customer_service/internal/service"
	desc "customer_service/pkg/grpc/customer_v1"
	"log"

	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	customer, err := s.customerService.New(ctx, req.GetName(), req.GetPhone(), req.GetPassport())

	if err != nil {
		return nil, err
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
