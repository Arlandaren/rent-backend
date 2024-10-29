package repository

import "time"

type CustomerDB struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Passport  string    `db:"passport"`
	CreatedAt time.Time `db:"created_at"`
}

type ApartmentDB struct {
	Id        int64     `db:"id"`
	Title     string    `db:"title"`
	Expenses  int64     `db:"expenses"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

type BookingDB struct {
	Id          int64     `db:"id"`
	ApartmentID int64     `db:"apartment_id"`
	DateStart   time.Time `db:"date_start"`
	DateEnd     time.Time `db:"date_end"`
	Price       int64     `db:"price"`
	CustomerID  int64     `db:"customer_id"`
	Status      string    `db:"status"`
	Comment     string    `db:"comment"`
	DateCreated time.Time `db:"date_created"`
}
