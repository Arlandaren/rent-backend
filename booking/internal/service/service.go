package service

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"service/internal/repository"
	"service/internal/shared/kafka"
	desc "service/pkg/grpc/booking_v1"
	"strconv"
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

func (s *Service) New(ctx context.Context, req *desc.NewBookingRequest) (*desc.Booking, error) {
	booking, err := s.repo.New(ctx, req)
	if err != nil {
		return nil, err
	}

	event := map[string]string{
		"id":           strconv.Itoa(int(booking.Id)),
		"apartment_id": strconv.Itoa(int(booking.ApartmentId)),
		"date_start":   booking.DateStart.String(),
		"date_end":     booking.DateEnd.String(),
		"price":        strconv.Itoa(int(booking.Price)),
		"customer_id":  strconv.Itoa(int(booking.CustomerId)),
		"status":       booking.Status,
		"date_created": booking.DateCreated.String(),
		"comment":      booking.Comment,
	}

	msg, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	err = s.producer.ProduceMessage("booking_created", msg)
	if err != nil {
		return nil, err
	}

	log.Println("Booking created")

	return booking, nil
}

func (s *Service) Begin(ctx context.Context, req *desc.BeginBookingRequest) (*desc.BeginBookingResponse, error) {
	date, err := s.repo.Begin(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Println("Booking begun")

	event := map[string]string{
		"id":         strconv.Itoa(int(req.Id)),
		"date_start": date.String(),
	}

	msg, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	err = s.producer.ProduceMessage("booking_begin", msg)
	if err != nil {
		return nil, err
	}

	response := &desc.BeginBookingResponse{
		Id:        req.Id,
		DateStart: timestamppb.New(date),
	}

	return response, nil

}

func (s *Service) Finish(ctx context.Context, req *desc.FinishBookingRequest) (*desc.FinishBookingResponse, error) {
	date, err := s.repo.Finish(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Println("Booking finished")

	event := map[string]string{
		"id":       strconv.Itoa(int(req.Id)),
		"date_end": date.String(),
	}

	msg, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	err = s.producer.ProduceMessage("booking_finished", msg)
	if err != nil {
		return nil, err
	}

	response := &desc.FinishBookingResponse{
		Id:      req.Id,
		DateEnd: timestamppb.New(date),
	}

	return response, nil
}

func (s *Service) Cancel(ctx context.Context, req *desc.CancelBookingRequest) (*desc.CancelBookingResponse, error) {
	status, err := s.repo.Cancel(ctx, req)
	if err != nil {
		return nil, err
	}

	event := map[string]string{
		"id":     strconv.Itoa(int(req.Id)),
		"status": status,
	}

	msg, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	err = s.producer.ProduceMessage("booking_cancelled", msg)
	if err != nil {
		return nil, err
	}

	log.Println("Booking cancelled")

	return &desc.CancelBookingResponse{Status: status}, nil

}
