package business

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type RegisteredAgent struct {
	Id        id.Id                   `json:"id,omitempty"`
	Business  Business                `json:"business,omitempty"`
	StartDate rfc3339date.Rfc3339Date `json:"start-date,omitempty"`
	EndDate   rfc3339date.Rfc3339Date `json:"end-date,omitempty"`
}
