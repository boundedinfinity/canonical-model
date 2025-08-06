package banking

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/currency"
)

type StatementTransaction struct {
	Amount currency.Amount         `json:"amount,omitempty"`
	Date   rfc3339date.Rfc3339Date `json:"date,omitempty"`
	Memo   string                  `json:"memo,omitempty"`
}
