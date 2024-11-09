package service

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"service/internal/shared/entities"
)

func (s *Service) ProcessBookingCreated(data *entities.BookingCreatedEvent) error {
	log.Println("event booking_created")

	ctx := context.Background()

	err := s.repo.NewBooking(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())

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

func (s *Service) ProcessBookingCancelled(message *sarama.ConsumerMessage) {

	// Здесь добавь
}

func (s *Service) ProcessApartmentCreated(message *sarama.ConsumerMessage) {

	// Здесь добавь
}

func (s *Service) ProcessApartmentRemoved(message *sarama.ConsumerMessage) {

	// Здесь добавь
}

func (s *Service) ProcessApartmentUpdated(message *sarama.ConsumerMessage) {

	// Здесь добавь
}

func (s *Service) ProcessCustomerCreated(message *sarama.ConsumerMessage) {
	fmt.Println("event customer_created")

	//s.repo.
	// Здесь добавь
}

func (s *Service) ProcessCustomerRemoved(message *sarama.ConsumerMessage) {

	// Здесь добавь
}

func (s *Service) ProcessCustomerUpdated(message *sarama.ConsumerMessage) {

	// Здесь добавь
}
