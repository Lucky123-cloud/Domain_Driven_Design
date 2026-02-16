package coffeeco

import uuid "github.com/google/uuid"

type CoffeeLover struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	EmailAddress string
}
