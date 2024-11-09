package repository

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"service/internal/shared/entities"
	"time"
)

func (r *Repository) NewBooking(ctx context.Context, booking *entities.BookingCreatedEvent) error {
	txWrapper, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if rbErr := txWrapper.Rollback(ctx); rbErr != nil {
				log.Printf("Ошибка при откате транзакции: %v", rbErr)
			}
		} else {
			if commitErr := txWrapper.Commit(ctx); commitErr != nil {
				err = commitErr
			}
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
		return fmt.Errorf("Ошибка вставки бронирования: %w", err)
	}

	_, err = txWrapper.Exec(ctx,
		"UPDATE apartment SET status = 'active' WHERE id = $1", booking.ApartmentID)
	if err != nil {
		return fmt.Errorf("Ошибка обновления статуса апартамента: %w", err)
	}

	return nil
}

func (r *Repository) NewApartment(ctx context.Context, apartment *entities.ApartmentCreatedEvent) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO apartment
		(id, address, status, date_created)
		VALUES ($1, $2, $3, $4)`,
		apartment.ID,
		apartment.Address,
		apartment.Status,
		time.Unix(apartment.DateCreated, 0),
	)
	if err != nil {
		return fmt.Errorf("Ошибка вставки апартамента: %w", err)
	}

	return nil
}
