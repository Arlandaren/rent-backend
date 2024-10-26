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

func (c *CustomerRepository) Remove(ctx context.Context, id int) error {
	_, err := c.db.Exec(ctx, "DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		log.Println("Failed to remove customer: ", err)
		return err
	}
	log.Println("RemoveCustomer")
	return nil
}

func (c *CustomerRepository) Update(ctx context.Context, id int, name, phone, passport string) error {
	_, err := c.db.Exec(ctx, "UPDATE customers SET name = $1, phone = $2, passport = $3 WHERE id = $4", name, phone, passport, id)
	if err != nil {
		return err
	}
	log.Println("UpdateCustomer")
	return nil
}