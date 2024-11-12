package repository

import (
	"context"
	"github.com/Arlandaren/pgxWrappy/postgres"
	"log"

	//"service/internal/shared/storage/postgres"

	desc "service/pkg/grpc/aggregator_v1"
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

func (r *Repository) CustomerGet(ctx context.Context, request *desc.CustomerGetRequest) (*CustomerDB, error) {
	var customer CustomerDB
	err := r.db.Get(ctx, &customer, "SELECT * FROM customers WHERE id = $1", request.Id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *Repository) CustomerList(ctx context.Context, request *desc.CustomerListRequest) ([]*CustomerDB, error) {
	var customers []*CustomerDB
	var err error

	if request.Cursor == 0 {
		err = r.db.Select(ctx, &customers, "SELECT * FROM customers ORDER BY id DESC LIMIT $1", request.Limit)
	} else {
		err = r.db.Select(ctx, &customers, "SELECT * FROM customers WHERE id < $1 ORDER BY id DESC LIMIT $2", request.Cursor, request.Limit)
	}

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *Repository) ApartmentGet(ctx context.Context, request *desc.ApartmentGetRequest) (*ApartmentDB, error) {
	var apartment ApartmentDB
	err := r.db.Get(ctx, &apartment, "SELECT * FROM apartment WHERE id = $1", request.Id)
	if err != nil {
		return nil, err
	}

	return &apartment, nil
}

func (r *Repository) ApartmentList(ctx context.Context, request *desc.ApartmentListRequest) ([]*ApartmentDB, error) {
	var apartment []*ApartmentDB

	var err error

	if request.Cursor == 0 {
		err = r.db.Select(ctx, &apartment, "SELECT * FROM apartment ORDER BY id DESC LIMIT $1", request.Limit)
	} else {
		err = r.db.Select(ctx, &apartment, "SELECT * FROM apartment WHERE id < $1 ORDER BY id DESC LIMIT $2", request.Cursor, request.Limit)
	}

	if err != nil {
		return nil, err
	}

	return apartment, nil
}

func (r *Repository) BookingGet(ctx context.Context, request *desc.BookGetRequest) (*BookingDB, error) {
	var booking BookingDB
	err := r.db.Get(ctx, &booking, "SELECT * FROM booking WHERE id = $1", request.Id)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *Repository) BookingList(ctx context.Context, request *desc.BookListRequest) ([]*BookingDB, error) {
	var bookings []*BookingDB

	var err error

	if request.Cursor == 0 {
		err = r.db.Select(ctx, &bookings, "SELECT * FROM booking ORDER BY id DESC LIMIT $1", request.Limit)
	} else {
		err = r.db.Select(ctx, &bookings, "SELECT * FROM booking WHERE id < $1 ORDER BY id DESC LIMIT $2", request.Cursor, request.Limit)
	}

	if err != nil {
		return nil, err
	}

	return bookings, nil
}
