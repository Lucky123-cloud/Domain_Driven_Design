package purchase

import (
	"context"
	"errors"
	"fmt"
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
	PaymentMeans       payment.Means
	TimeofPurchase     time.Time
	CardToken          *string
}

type CardChargeService interface {
	ChargeCard(ctx context.Context, amount money.Money, cardToken string) error
}

type Service struct {
	cardService  CardChargeService
	purchaseRepo Repository
}

func (s *Service) CompletePurchase(ctx context.Context, purchase *Purchase) error {
	if purchase == nil {
		return errors.New("purchase cannot be nil")
	}

	if err := purchase.validateAndEnrich(); err != nil {
		return fmt.Errorf("purchase validation failed: %w", err)
	}

	switch purchase.PaymentMeans {

	case payment.MEANS_CARD:
		if purchase.cardToken == nil {
			return errors.New("card token is required for card payment")
		}

		if err := s.cardService.ChargeCard(
			ctx,
			purchase.total,
			*purchase.cardToken,
		); err != nil {
			return fmt.Errorf("card charge failed: %w", err)
		}

	case payment.MEANS_CASH:
		// Cash requires no external processing.

	default:
		return errors.New("unknown payment type")
	}

	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return fmt.Errorf("failed to store purchase: %w", err)
	}

	return nil
}
