package business

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

type RegisteredAgent struct {
	Id        ider.Id                 `json:"id,omitempty"`
	Business  Business                `json:"business,omitempty"`
	StartDate rfc3339date.Rfc3339Date `json:"start-date,omitempty"`
	EndDate   rfc3339date.Rfc3339Date `json:"end-date,omitempty"`
}
