package coffeeco

import (
	"errors"

	"github.com/Rhymond/go-money"
)

type Product struct {
	ItemName  string
	BasePrice money.Money
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) == 0 {
		return errors.New("purchase must consist of atleast one product")
	}

	p.total = *money.New(0, "USD")
	for _, v := range p.ProductsToPurchase {
		newTotal, _ := p.total.Add(&v.BasePrice)
		p.total = *newTotal
	}

	return nil
}
