package bookkeeping

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/currency"
	"github.com/boundedinfinity/schema/idiomatic/digital_document"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Transaction struct {
	Id          id.Id                         `json:"id,omitempty"`
	Movement    []Movement                    `json:"movement,omitempty"`
	Amount      float32                       `json:"amount,omitempty"`
	Description string                        `json:"description,omitempty"`
	CheckNumber int                           `json:"check-number,omitempty"`
	Attachments []digital_document.DocumentV1 `json:"attachements,omitempty"`
}

type Movement struct {
	Account Account                 `json:"account,omitempty"`
	Amount  currency.Amount         `json:"amount,omitempty"`
	Date    rfc3339date.Rfc3339Date `json:"date,omitempty"`
	Memo    string                  `json:"memo,omitempty"`
}
