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

	event := map[string]string{
		"id":         fmt.Sprintf("%d", apartment.ID),
		"title":      apartment.Title,
		"expenses":   fmt.Sprintf("%d", apartment.Expenses),
		"created_at": apartment.CreatedAt.String(),
	}

	message, err := json.Marshal(event)
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
