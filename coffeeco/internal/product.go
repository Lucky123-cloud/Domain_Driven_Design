package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Product struct {
	ItemName  string
	BasePrice money.Money
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) == 0 {
		return errors.New("purchase must contain at least one product")
	}

	total := money.New(0, "USD")

	for _, product := range p.ProductsToPurchase {
		newTotal, err := total.Add(&product.BasePrice)
		if err != nil {
			return fmt.Errorf("failed to calculate total: %w", err)
		}
		total = newTotal
	}

	if total.IsZero() {
		return errors.New("invalid purchase: total amount cannot be zero")
	}

	p.total = *total
	p.id = uuid.New()
	p.timeOfPurchase = time.Now()

	return nil
}
