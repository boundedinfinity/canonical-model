package payment_card

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/payment"
)

type DebitCard struct {
	Card
}

var _ payment.Payment = &DebitCard{}

func (_ DebitCard) Kind() payment.Kind {
	return payment.Kinds.DebitCard
}

func (this DebitCard) String() string {
	return "debit card: " + this.Card.Number.Last4Digits()
}
