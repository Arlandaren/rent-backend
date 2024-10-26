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
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Expenses  int64     `json:"expenses"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewRepository(wrapper *postgres.Wrapper) *Repository {
	log.Println("NewRepository")
	return &Repository{
		db: wrapper,
	}
}

func (r *Repository) New(ctx context.Context, title string, expenses int64) (*Apartment, error) {
	var apartment Apartment
	err := r.db.QueryRow(
		ctx,
		"INSERT INTO apartment (title, expenses) VALUES ($1, $2) RETURNING id, status,created_at",
		title, expenses,
	).Scan(&apartment.ID, &apartment.Status, &apartment.CreatedAt)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	apartment.Title = title
	apartment.Expenses = expenses

	return &apartment, nil
}

func (r *Repository) Remove(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, "DELETE FROM apartment WHERE id = $1", id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, id, expenses int64, status, title string) (*Apartment, error) {
	var apartment Apartment

	err := r.db.QueryRow(ctx, `
        UPDATE apartment
        SET title = $1, expenses = $2, status = $3
        WHERE id = $4
        RETURNING id, title, expenses, status, created_at
    `, title, expenses, status, id).Scan(
		&apartment.ID, &apartment.Title, &apartment.Expenses, &apartment.Status, &apartment.CreatedAt,
	)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	return &apartment, nil
}
