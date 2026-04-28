package payment_card

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/payment"
)

type DebitCard struct {
	// CardModel
	// BankingAccount account.AccountModel `json:"banking-account,omitempty"`
}

var _ payment.Payment = &DebitCard{}

func (_ DebitCard) Kind() payment.Kind {
	return payment.Kinds.DebitCard
}
