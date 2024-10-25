package service

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type CustomerService struct {
	db *pgxpool.Pool
}

func NewCustomerService(db *pgxpool.Pool) *CustomerService {
	log.Println("NewCustomerService")
	return &CustomerService{
		db: db,
	}
}

func (s *CustomerService) NewCustomer(name, phone, passport string) (string, error) {

	log.Println("NewCustomer")

	return fmt.Sprintf("NewCustomer: %s, %s, %s", name, phone, passport), nil
}
