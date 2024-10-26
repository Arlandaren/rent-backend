package service

import (
	"context"
	"customer_service/internal/repository"
	"log"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	log.Println("NewCustomerService")
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) New(ctx context.Context, name, phone, passport string) (*repository.Customer, error) {
	customer, err := s.repo.New(ctx, name, phone, passport)
	if err != nil {
		return nil, err
	}
	log.Println("NewCustomer")
	return customer, nil
}
