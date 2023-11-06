package banking

import (
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/business"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Bank struct {
	Id       id.Id             `json:"id,omitempty"`
	Business business.Business `json:"business,omitempty"`
	Audit    audit.Audit       `json:"audit,omitempty"`
}
