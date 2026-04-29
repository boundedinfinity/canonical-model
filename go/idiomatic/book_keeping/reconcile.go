package book_keeping

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/document"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

type Reconcile struct {
	Id              ider.Id                 `json:"id,omitempty"`
	StatementDate   rfc3339date.Rfc3339Date `json:"statement-date,omitempty"`
	StartingBalance float32                 `json:"starting-balance,omitempty"`
	EndingBalance   float32                 `json:"ending-balance,omitempty"`
	Description     string                  `json:"description,omitempty"`
	Memo            string                  `json:"memo,omitempty"`
	Attachment      document.Document       `json:"attachement,omitempty"`
}
