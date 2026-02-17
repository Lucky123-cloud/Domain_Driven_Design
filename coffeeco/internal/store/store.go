package store

import (
	coffeeco "github.com/Lucky123-cloud/Domain_Driven_Design/internal"

	"github.com/google/uuid"
)

type Store struct {
	ID             uuid.UUID
	Location       string
	ProductForSale []coffeeco.Product
}
