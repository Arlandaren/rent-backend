package service

import "fmt"

func (s *Service) ProcessBookingCreated(data *BookingCreatedEvent) error {
	fmt.Println(data)

	//взаимодействие с базой данных

	return nil
}
