package repository

import (
	"context"
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"service/internal/shared/storage/postgres"
	desc "service/pkg/grpc/booking_v1"
	"time"
)

type Repository struct {
	db *postgres.Wrapper
}

func NewRepository(wrapper *postgres.Wrapper) *Repository {
	log.Println("NewRepository")
	return &Repository{
		db: wrapper,
	}
}

func (r *Repository) New(ctx context.Context, req *desc.NewBookingRequest) (*desc.Booking, error) {
	var booking desc.Booking
	var dateStart, dateEnd, dateCreated sql.NullTime
	err := r.db.QueryRow(ctx,
		"INSERT INTO booking (apartment_id, price, customer_id, comment) VALUES ($1, $2, $3, $4) RETURNING id, date_start, date_end, status, date_created",
		req.ApartmentId, req.Price, req.CustomerId, req.Comment).Scan(&booking.Id, &dateStart, &dateEnd, &booking.Status, &dateCreated)

	if err != nil {
		return nil, err
	}

	booking.DateStart = timestamppb.New(dateStart.Time)
	booking.DateEnd = timestamppb.New(dateEnd.Time)
	booking.DateCreated = timestamppb.New(dateCreated.Time)

	booking.ApartmentId = req.ApartmentId
	booking.Price = req.Price
	booking.CustomerId = req.CustomerId
	booking.Comment = req.Comment

	return &booking, nil
}

func (r *Repository) Begin(ctx context.Context, req *desc.BeginBookingRequest) (time.Time, error) {
	var dateStart time.Time
	err := r.db.QueryRow(ctx,
		"UPDATE booking SET date_start = $1, status = 'active' WHERE id = $2 AND  status = 'pending' RETURNING date_start",
		time.Now(), req.Id).Scan(&dateStart)

	if err != nil {
		return time.Time{}, err
	}

	return dateStart, nil
}

func (r *Repository) Finish(ctx context.Context, req *desc.FinishBookingRequest) (time.Time, error) {
	var dateEnd time.Time
	err := r.db.QueryRow(ctx,
		"UPDATE booking SET date_end = $1, status = 'finished' WHERE id = $2 AND  status = 'active' RETURNING date_end",
		time.Now(), req.Id).Scan(&dateEnd)

	if err != nil {
		return time.Time{}, err
	}

	return dateEnd, nil
}
