package repository

import (
	"service/internal/shared/storage/postgres"

	"log"
	"time"

	"context"
)

type Repository struct {
	db *postgres.Wrapper
}

type Customer struct {
	ID        int64
	Name      string
	Phone     string
	Passport  string
	CreatedAt time.Time
}

func NewRepository(wrapper *postgres.Wrapper) *Repository {
	log.Println("NewCustomerRepository")
	return &Repository{
		db: wrapper,
	}
}

func (c *Repository) New(ctx context.Context, name, phone, passport string) (*Customer, error) {
	var id int64
	var createdAt time.Time
	query := "INSERT INTO customers (name, phone, passport) VALUES ($1, $2, $3) RETURNING id, created_at"

	err := c.db.QueryRow(ctx, query, name, phone, passport).Scan(&id, &createdAt)
	if err != nil {
		return nil, err
	}

	customer := &Customer{
		ID:        id,
		Name:      name,
		Phone:     phone,
		Passport:  passport,
		CreatedAt: createdAt,
	}
	log.Println("NewCustomer")

	return customer, nil
}

func (c *Repository) Remove(ctx context.Context, id int64) error {
	_, err := c.db.Exec(ctx, "DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		log.Println("Failed to remove customer: ", err)
		return err
	}
	log.Println("RemoveCustomer")
	return nil
}

func (c *Repository) Update(ctx context.Context, id int64, name, phone, passport string) error {
	_, err := c.db.Exec(ctx, "UPDATE customers SET name = $1, phone = $2, passport = $3 WHERE id = $4", name, phone, passport, id)
	if err != nil {
		return err
	}
	log.Println("UpdateCustomer")
	return nil
}
