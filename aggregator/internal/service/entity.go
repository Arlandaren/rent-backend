package service

type BookingCreatedEvent struct {
	ApartmentID string `json:"apartment_id"`
	Comment     string `json:"comment"`
	CustomerID  string `json:"customer_id"`
	DateCreated string `json:"date_created"`
	DateEnd     string `json:"date_end"`
	DateStart   string `json:"date_start"`
	ID          string `json:"id"`
	Price       string `json:"price"`
	Status      string `json:"status"`
}
