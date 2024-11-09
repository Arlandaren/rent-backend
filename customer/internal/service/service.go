package service

import (
	"context"
	"service/internal/repository"
	"service/internal/shared/entities"

	"service/internal/shared/kafka"

	"encoding/json"

	"log"
)

type Service struct {
	repo     *repository.Repository
	producer *kafka.Producer
}

func NewService(repo *repository.Repository, producer *kafka.Producer) *Service {
	log.Println("Service")
	if producer == nil {
		log.Fatal("producer is nil")
	}
	return &Service{
		repo:     repo,
		producer: producer,
	}
}

func (s *Service) New(ctx context.Context, name, phone, passport string) (*repository.Customer, error) {
	customer, err := s.repo.New(ctx, name, phone, passport)
	if err != nil {
		return nil, err
	}

	event := entities.CustomerCreatedEvent{
		ID:        customer.ID,
		Name:      customer.Name,
		Phone:     customer.Phone,
		Passport:  customer.Passport,
		CreatedAt: customer.CreatedAt.Unix(),
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v", err)
		return nil, err
	}

	err = s.producer.ProduceMessage("customer_created", message)
	if err != nil {
		log.Printf("Failed to produce message to Kafka: %v", err)
		return nil, err
	}

	log.Println("NewCustomer")

	return customer, nil
}

func (s *Service) Remove(ctx context.Context, id int64) error {
	err := s.repo.Remove(ctx, id)
	if err != nil {
		log.Printf("Failed to remove customer: %v", err)
		return err
	}

	event := entities.CustomerRemovedEvent{
		ID: id,
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v", err)
		return err
	}

	err = s.producer.ProduceMessage("customer_removed", message)
	if err != nil {
		return err
	}

	log.Println("RemoveCustomer")

	return nil
}

func (s *Service) Update(ctx context.Context, id int64, name, phone, passport string) error {
	err := s.repo.Update(ctx, id, name, phone, passport)
	if err != nil {
		log.Printf("Failed to update customer: %v", err)
		return err
	}

	event := entities.CustomerUpdatedEvent{
		ID:       id,
		Name:     name,
		Phone:    phone,
		Passport: passport,
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v", err)
		return err
	}

	err = s.producer.ProduceMessage("customer_updated", message)
	if err != nil {
		return err
	}

	log.Println("UpdateCustomer")

	return nil
}
