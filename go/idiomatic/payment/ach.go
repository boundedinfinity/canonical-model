package payment

import (
	"fmt"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/banking/account"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/business"
)

type AutomatedClearingHouse struct {
	Account account.Account   `json:"bank-account,omitempty"`
	Bank    business.Business `json:"bank,omitempty"`
}

var _ Payment = &AutomatedClearingHouse{}

func (_ AutomatedClearingHouse) Kind() Kind {
	return Kinds.AutomatedClearningHouse
}

func (this AutomatedClearingHouse) String() string {
	return fmt.Sprintf("ACH: %s : %s", this.Account.Last4Digits(), this.Bank)
}
