package bookkeeping

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/digital_document"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Reconcile struct {
	Id              id.Id                        `json:"id,omitempty"`
	StatementDate   rfc3339date.Rfc3339Date      `json:"statement-date,omitempty"`
	StartingBalance float32                      `json:"starting-balance,omitempty"`
	EndingBalance   float32                      `json:"ending-balance,omitempty"`
	Description     string                       `json:"description,omitempty"`
	Memo            string                       `json:"memo,omitempty"`
	Attachment      *digital_document.DocumentV1 `json:"attachement,omitempty"`
}
