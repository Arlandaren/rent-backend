package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"service/internal/repository"
	"service/internal/shared/kafka"
	desc "service/pkg/grpc/aggregator_v1"
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

func (s *Service) CustomerGet(ctx context.Context, request *desc.CustomerGetRequest) (*desc.CustomerGetResponse, error) {
	log.Println("customer get")
	customer, err := s.repo.CustomerGet(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := &desc.Customer{
		Id:        customer.Id,
		Name:      customer.Name,
		Phone:     customer.Phone,
		Passport:  customer.Passport,
		CreatedAt: timestamppb.New(customer.CreatedAt),
	}

	return &desc.CustomerGetResponse{Customer: response}, nil
}

func (s *Service) CustomerList(ctx context.Context, request *desc.CustomerListRequest) (*desc.CustomerListResponse, error) {
	customersDB, err := s.repo.CustomerList(ctx, request)
	if err != nil {
		return nil, err
	}

	customers := make([]*desc.Customer, 0, len(customersDB))
	for _, cdb := range customersDB {
		customer := &desc.Customer{
			Id:        cdb.Id,
			Name:      cdb.Name,
			Phone:     cdb.Phone,
			Passport:  cdb.Passport,
			CreatedAt: timestamppb.New(cdb.CreatedAt),
		}
		customers = append(customers, customer)
	}

	var nextCursor int64
	if len(customersDB) > 0 {
		nextCursor = customersDB[len(customersDB)-1].Id
	} else {
		nextCursor = 0
	}

	return &desc.CustomerListResponse{
		Customers:  customers,
		NextCursor: nextCursor,
	}, nil
}

func (s *Service) ApartmentGet(ctx context.Context, request *desc.ApartmentGetRequest) (*desc.ApartmentGetResponse, error) {
	log.Println("apartment get")
	apartment, err := s.repo.ApartmentGet(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := &desc.Apartment{
		Id:        apartment.Id,
		Title:     apartment.Title,
		Expenses:  apartment.Expenses,
		Status:    apartment.Status,
		CreatedAt: timestamppb.New(apartment.CreatedAt),
	}

	return &desc.ApartmentGetResponse{Apartment: response}, nil
}

func (s *Service) ApartmentList(ctx context.Context, request *desc.ApartmentListRequest) (*desc.ApartmentListResponse, error) {
	apartmentsDB, err := s.repo.ApartmentList(ctx, request)
	if err != nil {
		return nil, err
	}

	apartments := make([]*desc.Apartment, 0, len(apartmentsDB))
	for _, adb := range apartmentsDB {
		apartment := &desc.Apartment{
			Id:        adb.Id,
			Title:     adb.Title,
			Expenses:  adb.Expenses,
			Status:    adb.Status,
			CreatedAt: timestamppb.New(adb.CreatedAt),
		}
		apartments = append(apartments, apartment)
	}

	var nextCursor int64
	if len(apartmentsDB) > 0 {
		nextCursor = apartmentsDB[len(apartmentsDB)-1].Id
	} else {
		nextCursor = 0
	}

	return &desc.ApartmentListResponse{
		Apartments: apartments,
		NextCursor: nextCursor,
	}, nil
}

func (s *Service) BookingGet(ctx context.Context, request *desc.BookGetRequest) (*desc.BookGetResponse, error) {
	log.Println("Booking get")
	booking, err := s.repo.BookingGet(ctx, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := &desc.Booking{
		Id:          booking.Id,
		ApartmentId: booking.ApartmentID,
		DateStart:   timestamppb.New(booking.DateStart),
		DateEnd:     timestamppb.New(booking.DateEnd),
		Price:       booking.Price,
		CustomerId:  booking.CustomerID,
		Status:      booking.Status,
		Comment:     booking.Comment,
		DateCreated: timestamppb.New(booking.DateCreated),
	}

	return &desc.BookGetResponse{Booking: response}, nil
}

func (s *Service) BookingList(ctx context.Context, request *desc.BookListRequest) (*desc.BookListResponse, error) {
	bookingsDB, err := s.repo.BookingList(ctx, request)
	if err != nil {
		return nil, err
	}

	bookings := make([]*desc.Booking, 0, len(bookingsDB))
	for _, bdb := range bookingsDB {
		booking := &desc.Booking{
			Id:          bdb.Id,
			ApartmentId: bdb.ApartmentID,
			DateStart:   timestamppb.New(bdb.DateStart),
			DateEnd:     timestamppb.New(bdb.DateEnd),
			Price:       bdb.Price,
			CustomerId:  bdb.CustomerID,
			Status:      bdb.Status,
			Comment:     bdb.Comment,
			DateCreated: timestamppb.New(bdb.DateCreated),
		}
		bookings = append(bookings, booking)
	}

	var nextCursor int64
	if len(bookingsDB) > 0 {
		nextCursor = bookingsDB[len(bookingsDB)-1].Id
	} else {
		nextCursor = 0
	}

	return &desc.BookListResponse{
		Bookings:   bookings,
		NextCursor: nextCursor,
	}, nil
}
