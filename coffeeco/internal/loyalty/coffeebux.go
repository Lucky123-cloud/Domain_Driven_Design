package loyalty

import (
	coffeeco "github.com/Lucky123-cloud/Domain_Driven_Design/internal"
	"github.com/Lucky123-cloud/Domain_Driven_Design/internal/store"
	"github.com/google/uuid"
)

type Coffeebux struct {
	ID                                    uuid.UUID
	store                                 store.Store
	coffeeLover                           coffeeco.CoffeeLover
	FreeDrinksAvailable                   int
	RemainingDrinkPurchasesUntilFreeDrink int
}
