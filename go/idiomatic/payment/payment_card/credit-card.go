package payment_card

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/payment"
)

type CreditCard struct {
	CardModel
}

var _ payment.Payment = &CreditCard{}

func (_ CreditCard) Kind() payment.Kind {
	return payment.Kinds.CreditCard
}
