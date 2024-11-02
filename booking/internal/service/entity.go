package service

type bookingCreatedEvent struct {
	ID          int64  `json:"id"`
	ApartmentID int64  `json:"apartment_id"`
	DateEnd     int64  `json:"date_end"`   // UNIX timestamp в секундах
	DateStart   int64  `json:"date_start"` // UNIX timestamp в секундах
	Price       int64  `json:"price"`
	CustomerID  int64  `json:"customer_id"`
	Status      string `json:"status"`
	DateCreated int64  `json:"date_created"` // UNIX timestamp в секундах
	Comment     string `json:"comment"`
}

type bookingBeganEvent struct {
	ID        int64 `json:"id"`
	DateStart int64 `json:"date_start"`
}

type bookingUpdatedEvent struct {
	ID          int64  `json:"id"`
	ApartmentID int64  `json:"apartment_id"`
	Price       int64  `json:"price"`
	CustomerID  int64  `json:"customer_id"`
	Comment     string `json:"comment"`
}

type bookingFinishedEvent struct {
	ID     int
	Status string
}
