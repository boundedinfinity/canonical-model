package banking

import (
	"github.com/boundedinfinity/schema/idiomatic/business"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Bank struct {
	Id       id.Id             `json:"id,omitempty"`
	Business business.Business `json:"business,omitempty"`
}
