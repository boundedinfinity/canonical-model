package currency

import "github.com/boundedinfinity/schema/idiomatic/id"

type CurrencyAmount struct {
	Id      id.Id        `json:"id,omitempty"`
	Type    CurrencyType `json:"type,omitempty"`
	Account float32      `json:"account,omitempty"`
}
