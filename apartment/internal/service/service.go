package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"service/internal/repository"
	"service/internal/shared/kafka"
)

type Service struct {
	repo     *repository.Repository
	producer *kafka.Producer
}

func NewService(repo *repository.Repository, producer *kafka.Producer) *Service {
	log.Println("NewService")
	return &Service{
		repo:     repo,
		producer: producer,
	}
}

func (s *Service) New(ctx context.Context, title string, expenses int64) (*repository.Apartment, error) {

	apartment, err := s.repo.New(ctx, title, expenses)
	if err != nil {
		return nil, err
	}
	log.Println("NewApartment")

	message, err := json.Marshal(apartment)
	if err != nil {
		return nil, err
	}

	err = s.producer.ProduceMessage("apartment_created", message)
	if err != nil {
		return nil, err
	}

	return apartment, nil
}

func (s *Service) Remove(ctx context.Context, id int64) error {
	if err := s.repo.Remove(ctx, id); err != nil {
		log.Printf("Failed to remove apartment: %v", err)
		return err
	}

	err := s.producer.ProduceMessage("apartment_removed", []byte(fmt.Sprintf(`{"id": "%d"}`, id)))
	if err != nil {
		return err
	}
	log.Println("RemoveApartment")
	return nil
}

func (s *Service) Update(ctx context.Context, id, expenses int64, status, title string) (*repository.Apartment, error) {

	apartment, err := s.repo.Update(ctx, id, expenses, status, title)
	if err != nil {
		return nil, err
	}

	message, err := json.Marshal(apartment)
	if err != nil {
		return nil, err
	}

	err = s.producer.ProduceMessage("apartment_updated", message)
	if err != nil {
		return nil, err
	}

	log.Println("UpdateApartment")

	return apartment, nil
}
