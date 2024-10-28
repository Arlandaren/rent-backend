package service

import (
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

func (s Service) ProcessBookingCreated(data *BookingCreatedEvent) error {
	fmt.Println(data)

	//взаимодействие с базой данных

	return nil
}
