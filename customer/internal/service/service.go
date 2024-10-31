package service

import (
	"context"
	"service/internal/repository"

	"service/internal/shared/kafka"

	"encoding/json"

	"log"
	"strconv"
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

	event := map[string]string{
		"id":       strconv.Itoa(customer.ID),
		"name":     name,
		"phone":    phone,
		"passport": passport,
		"created":  strconv.FormatInt(customer.CreatedAt.Unix(), 10),
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

func (s *Service) Remove(ctx context.Context, id int) error {
	err := s.repo.Remove(ctx, id)
	if err != nil {
		log.Printf("Failed to remove customer: %v", err)
		return err
	}

	event := map[string]string{
		"id": strconv.Itoa(id),
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

func (s *Service) Update(ctx context.Context, id int, name, phone, passport string) error {
	err := s.repo.Update(ctx, id, name, phone, passport)
	if err != nil {
		log.Printf("Failed to update customer: %v", err)
		return err
	}

	event := map[string]string{
		"id":       strconv.Itoa(id),
		"name":     name,
		"phone":    phone,
		"passport": passport,
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
