package service

import (
	"context"
	"encoding/json"
	"log"
	"service/internal/repository"
	"service/internal/shared/kafka"
	"strconv"
)

type Service struct {
	repo     *repository.Repository
	producer *kafka.Producer
}

func NewService(repo *repository.Repository, producer *kafka.Producer) *Service {
	log.Println("NewService")

	if producer == nil {
		log.Fatal("producer is nil")
	}

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

	event := map[string]string{
		"id":         strconv.FormatInt(apartment.ID, 10),
		"title":      apartment.Title,
		"expenses":   strconv.FormatInt(apartment.Expenses, 10),
		"status":     apartment.Status,
		"created_at": strconv.FormatInt(apartment.CreatedAt.Unix(), 10),
	}

	message, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	err = s.producer.ProduceMessage("apartment_created", message)
	if err != nil {
		return nil, err
	}

	log.Println("NewApartment")

	return apartment, nil
}

func (s *Service) Remove(ctx context.Context, id int64) error {
	if err := s.repo.Remove(ctx, id); err != nil {
		log.Printf("Failed to remove apartment: %v", err)
		return err
	}

	event := map[string]string{
		"id": strconv.FormatInt(id, 10),
	}

	msg, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = s.producer.ProduceMessage("apartment_removed", msg)
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

	event := map[string]string{
		"id":       strconv.FormatInt(apartment.ID, 10),
		"title":    apartment.Title,
		"status":   apartment.Status,
		"expenses": strconv.FormatInt(apartment.Expenses, 10),
	}

	message, err := json.Marshal(event)
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
