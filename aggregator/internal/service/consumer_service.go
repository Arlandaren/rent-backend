package service

import (
	"context"
	"fmt"
	"log"
	"service/internal/shared/entities"
)

func (s *Service) ProcessBookingCreated(ctx context.Context, data *entities.BookingCreatedEvent) error {
	log.Println("event booking_created")

	err := s.repo.NewBooking(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())

		return err
	}

	return nil
}

func (s *Service) ProcessBookingBegan(ctx context.Context, data *entities.BookingBeganEvent) error {
	log.Println("event booking_began")

	err := s.repo.BeginBooking(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) ProcessBookingUpdated(ctx context.Context, data *entities.BookingUpdatedEvent) error {
	fmt.Println("event booking_updated")

	err := s.repo.UpdateBooking(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) ProcessBookingFinished(ctx context.Context, data *entities.BookingFinishedEvent) error {
	log.Println(data)

	err := s.repo.FinishBooking(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) ProcessBookingCancelled(ctx context.Context, data *entities.BookingCancelledEvent) error {
	log.Println("event booking_cancelled")

	err := s.repo.CancelBooking(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s *Service) ProcessApartmentCreated(ctx context.Context, data *entities.ApartmentCreatedEvent) error {
	log.Println("event apartment_created")

	err := s.repo.NewApartment(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}

	return nil

}

func (s *Service) ProcessApartmentRemoved(ctx context.Context, data *entities.ApartmentRemovedEvent) error {
	log.Println("event apartment_removed")

	err := s.repo.RemoveApartment(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s *Service) ProcessApartmentUpdated(ctx context.Context, data *entities.ApartmentUpdatedEvent) error {
	log.Println("event apartment_updated")

	err := s.repo.UpdateApartment(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s *Service) ProcessCustomerCreated(ctx context.Context, data *entities.CustomerCreatedEvent) error {
	log.Println("event customer_created")

	err := s.repo.NewCustomer(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) ProcessCustomerRemoved(ctx context.Context, data *entities.CustomerRemovedEvent) error {
	log.Println("event customer_removed")

	err := s.repo.RemoveCustomer(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s *Service) ProcessCustomerUpdated(ctx context.Context, data *entities.CustomerUpdatedEvent) error {
	log.Println("event customer_updated")

	err := s.repo.UpdateCustomer(ctx, data)
	if err != nil {
		//TODO: This error msg have to be sent to the notification microservice
		fmt.Println(err.Error())
		return err
	}
	return nil
}
