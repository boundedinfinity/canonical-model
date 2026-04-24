package business

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/people"
	"github.com/boundedinfinity/rfc3339date"
)

type ForeignEntity struct {
	Id               ider.Id                 `json:"id,omitempty"`
	LegalName        string                  `json:"legal-name,omitempty"`
	TradeName        string                  `json:"trade-name,omitempty"`
	Purpose          string                  `json:"purpose,omitempty"`
	Business         Business                `json:"business,omitempty"`
	RegisteredAgents []RegisteredAgent       `json:"registered-agents,omitempty"`
	SignedBy         []people.Person         `json:"signed-by,omitempty"`
	StartDate        rfc3339date.Rfc3339Date `json:"start-date,omitempty"`
	EndDate          rfc3339date.Rfc3339Date `json:"end-date,omitempty"`
}
