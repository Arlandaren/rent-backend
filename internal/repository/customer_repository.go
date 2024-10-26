package repository

import (
	"customer_service/internal/shared/storage/postgres"
	"log"
	"time"

	"context"
)

type CustomerRepository struct {
	db *postgres.Wrapper
}

type Customer struct {
	ID        int
	Name      string
	Phone     string
	Passport  string
	CreatedAt time.Time
}

func NewCustomerRepository(wrapper *postgres.Wrapper) *CustomerRepository {
	log.Println("NewCustomerRepository")
	return &CustomerRepository{
		db: wrapper,
	}
}

func (c *CustomerRepository) New(ctx context.Context, name, phone, passport string) (*Customer, error) {
	var id int
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
