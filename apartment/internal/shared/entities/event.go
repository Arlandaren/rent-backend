package entities

type ApartmentCreatedEvent struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Expenses  int64  `json:"expenses"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

type ApartmentUpdatedEvent struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Expenses int64  `json:"expenses"`
}

type ApartmentRemovedEvent struct {
	ID int64 `json:"id"`
}
