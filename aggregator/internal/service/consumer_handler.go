package service

import (
	"context"
	"fmt"
	"log"
	"service/internal/shared/entities"
)

func (s *Service) ProcessBookingCreated(data *entities.BookingCreatedEvent) error {
	log.Println("event booking_created")

	ctx := context.Background()

	err := s.repo.NewBooking(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ProcessBookingBegan(data *entities.BookingBeganEvent) error {
	fmt.Println(data)

	//взаимодействие с базой данных

	return nil
}

func (s *Service) ProcessBookingUpdated(data *entities.BookingBeganEvent) error {
	fmt.Println(data)

	//взаимодействие с базой данных

	return nil
}

func (s *Service) ProcessBookingFinished(data *entities.BookingBeganEvent) error {
	fmt.Println(data)

	//взаимодействие с базой данных

	return nil
}
