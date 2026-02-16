package purchase

import (
	"time"

	coffeeco "github.com/Lucky123-cloud/Domain_Driven_Design/internal"
	"github.com/Lucky123-cloud/Domain_Driven_Design/internal/payment"
	"github.com/Lucky123-cloud/Domain_Driven_Design/internal/store"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Purchase struct {
	id                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	total              money.Money
	paymentMeans       payment.Means
	timeOfPurchase     time.Time
	CardToken          *string
}
