package insurance

import (
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/business"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type InsuranceCompany struct {
	Id       id.Id             `json:"id,omitempty"`
	Number   string            `json:"number,omitempty"`
	Business business.Business `json:"business,omitempty"`
	Audit    audit.Audit       `json:"audit,omitempty"`
}
