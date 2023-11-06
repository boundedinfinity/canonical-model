package vehicle

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Vehicle struct {
	Id             id.Id                       `json:"id,omitempty"`
	Vin            VehicleIdentificationNumber `json:"vin,omitempty"`
	Make           string                      `json:"make,omitempty"`
	Model          string                      `json:"model,omitempty"`
	BodyStyle      string                      `json:"body-style,omitempty"`
	FuelEfficiency FuelEfficiency              `json:"fuel-efficiency,omitempty"`
	DriveType      string                      `json:"drive-type,omitempty"`
	FeulType       string                      `json:"fuel-type,omitempty"`
	EngineType     string                      `json:"engine-type,omitempty"`
	Transmission   string                      `json:"transmission,omitempty"`
	YearModel      rfc3339date.Rfc3339Date     `json:"year-model,omitempty"`
	PurchaseDate   rfc3339date.Rfc3339Date     `json:"purchase-date,omitempty"`
}

type TransmissionType string
type FeulType string
type EngineType string
type FuelEfficiency struct {
	City    int `json:"city,omitempty"`
	Highway int `json:"highway,omitempty"`
}
