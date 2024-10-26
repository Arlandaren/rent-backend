package repository

import (
	"context"
	"log"
	"service/internal/shared/storage/postgres"
	"time"
)

type Repository struct {
	db *postgres.Wrapper
}

type Apartment struct {
	ID        int64
	Title     string
	Expenses  int64
	CreatedAt time.Time
}

func NewRepository(wrapper *postgres.Wrapper) *Repository {
	log.Println("NewRepository")
	return &Repository{
		db: wrapper,
	}
}

func (r *Repository) New(ctx context.Context, title string, expenses int64) (*Apartment, error) {
	var id int
	var createdAt time.Time
	err := r.db.QueryRow(
		ctx,
		"INSERT INTO apartment (title, expenses) VALUES ($1, $2) RETURNING id, created_at",
		title, expenses,
	).Scan(&id, &createdAt)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	apartment := &Apartment{
		ID:        int64(id),
		Title:     title,
		Expenses:  expenses,
		CreatedAt: createdAt,
	}

	return apartment, nil
}

func (r *Repository) Remove(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, "DELETE FROM apartment WHERE id = $1", id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	return nil
}
