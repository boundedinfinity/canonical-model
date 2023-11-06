package vehicle

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Vehicle struct {
	Id           id.Id                       `json:"id,omitempty"`
	Vin          VehicleIdentificationNumber `json:"vin,omitempty"`
	Make         string                      `json:"make,omitempty"`
	Model        string                      `json:"model,omitempty"`
	YearModel    rfc3339date.Rfc3339Date     `json:"year-model,omitempty"`
	PurchaseDate rfc3339date.Rfc3339Date     `json:"purchase-date,omitempty"`
}
