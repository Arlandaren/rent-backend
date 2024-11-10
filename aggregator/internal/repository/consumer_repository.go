package repository

import (
	"context"
	"fmt"
	"service/internal/shared/entities"
	"time"
)

func (r *Repository) NewBooking(ctx context.Context, booking *entities.BookingCreatedEvent) (err error) {
	txWrapper, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = txWrapper.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = txWrapper.Rollback(ctx)
		} else {
			err = txWrapper.Commit(ctx)
		}
	}()

	_, err = txWrapper.Exec(ctx,
		`INSERT INTO booking
        (id, apartment_id, date_start, date_end, price, customer_id, status, comment, date_created)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		booking.ID,
		booking.ApartmentID,
		time.Unix(booking.DateStart, 0),
		time.Unix(booking.DateEnd, 0),
		booking.Price,
		booking.CustomerID,
		booking.Status,
		booking.Comment,
		time.Unix(booking.DateCreated, 0),
	)
	if err != nil {
		return fmt.Errorf("error inserting booking: %w", err)
	}

	_, err = txWrapper.Exec(ctx,
		"UPDATE apartment SET status = 'active' WHERE id = $1", booking.ApartmentID)
	if err != nil {
		return fmt.Errorf("error updating status of apartment: %w", err)
	}

	return nil
}

func (r *Repository) NewCustomer(ctx context.Context, customer *entities.CustomerCreatedEvent) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO customers
		(id, name, phone, passport, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		customer.ID,
		customer.Name,
		customer.Phone,
		customer.Passport,
		time.Unix(customer.CreatedAt, 0),
	)
	if err != nil {
		return fmt.Errorf("error creating customer: %w", err)
	}

	return nil
}

func (r *Repository) RemoveCustomer(ctx context.Context, customer *entities.CustomerRemovedEvent) error {
	_, err := r.db.Exec(ctx,
		"DELETE FROM customers WHERE id = $1",
		customer.ID,
	)
	if err != nil {
		return fmt.Errorf("error deleting customer: %w", err)
	}

	return nil
}

func (r *Repository) UpdateCustomer(ctx context.Context, customer *entities.CustomerUpdatedEvent) error {
	_, err := r.db.Exec(ctx,
		"UPDATE customers SET name = $1, phone = $2, passport = $3 WHERE id = $4",
		customer.Name,
		customer.Phone,
		customer.Passport,
		customer.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating customer: %w", err)
	}
	return nil
}

func (r *Repository) NewApartment(ctx context.Context, apartment *entities.ApartmentCreatedEvent) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO apartment
		(id,title, expenses, status, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		apartment.ID,
		apartment.Title,
		apartment.Expenses,
		apartment.Status,
		time.Unix(apartment.CreatedAt, 0),
	)
	if err != nil {
		return fmt.Errorf("error creating apartment: %w", err)
	}

	return nil
}

func (r *Repository) RemoveApartment(ctx context.Context, apartment *entities.ApartmentRemovedEvent) error {
	_, err := r.db.Exec(ctx,
		"DELETE FROM apartment WHERE id = $1",
		apartment.ID,
	)
	if err != nil {
		return fmt.Errorf("error deleting apartment: %w", err)
	}

	return nil
}

func (r *Repository) UpdateApartment(ctx context.Context, apartment *entities.ApartmentUpdatedEvent) error {
	_, err := r.db.Exec(ctx,
		"UPDATE apartment SET title = $1, expenses = $2, status = $3 WHERE id = $4",
		apartment.Title,
		apartment.Expenses,
		apartment.Status,
		apartment.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating apartment: %w", err)
	}

	return nil
}

func (r *Repository) BeginBooking(ctx context.Context, booking *entities.BookingBeganEvent) error {
	_, err := r.db.Exec(ctx,
		"UPDATE booking SET date_start=$1,status = 'active' WHERE id = $2",
		time.Unix(booking.DateStart, 0),
		booking.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating status of booking: %w", err)
	}

	return nil
}

func (r *Repository) UpdateBooking(ctx context.Context, booking *entities.BookingUpdatedEvent) error {
	_, err := r.db.Exec(ctx,
		"UPDATE booking SET apartment_id=$1,price=$2, customer_id=$3, comment=$4 WHERE id = $5",
		booking.ApartmentID,
		booking.Price,
		booking.CustomerID,
		booking.Comment,
		booking.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating status of booking: %w", err)
	}
	return nil
}

func (r *Repository) FinishBooking(ctx context.Context, booking *entities.BookingFinishedEvent) (err error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	_, err = tx.Exec(ctx,
		"UPDATE booking SET date_end=$1, status='finished' WHERE id=$2 AND status='active'",
		time.Unix(booking.DateEnd, 0),
		booking.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating status of booking: %w", err)
	}

	var apartmentID int64
	err = tx.QueryRow(ctx, "SELECT apartment_id FROM booking WHERE id = $1", booking.ID).Scan(&apartmentID)
	if err != nil {
		return fmt.Errorf("error getting apartment id: %w", err)
	}

	_, err = tx.Exec(ctx,
		"UPDATE apartment SET status='inactive' WHERE id=$1",
		apartmentID,
	)
	if err != nil {
		return fmt.Errorf("error updating status of apartment: %w", err)
	}

	return nil
}

func (r *Repository) CancelBooking(ctx context.Context, booking *entities.BookingCancelledEvent) (err error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	_, err = tx.Exec(ctx,
		"UPDATE booking SET status = 'cancelled' WHERE id = $1",
		booking.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating status of booking: %w", err)
	}

	var apartmentID int64
	err = tx.QueryRow(ctx, "SELECT apartment_id FROM booking WHERE id = $1", booking.ID).Scan(&apartmentID)
	if err != nil {
		return fmt.Errorf("error getting apartment id: %w", err)
	}

	_, err = tx.Exec(ctx,
		"UPDATE apartment SET status = 'inactive' WHERE id = $1",
		apartmentID,
	)
	if err != nil {
		return fmt.Errorf("error updating status of apartment: %w", err)
	}

	return nil
}
