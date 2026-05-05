package payment_card

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/payment"
)

type CreditCard struct {
	Card
}

var _ payment.Payment = &CreditCard{}

func (_ CreditCard) Kind() payment.Kind {
	return payment.Kinds.CreditCard
}
