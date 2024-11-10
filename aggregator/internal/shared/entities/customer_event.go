package entities

type CustomerCreatedEvent struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Passport  string `json:"passport"`
	CreatedAt int64  `json:"created_at"`
}

type CustomerUpdatedEvent struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Passport string `json:"passport"`
}

type CustomerRemovedEvent struct {
	ID int64 `json:"id"`
}
