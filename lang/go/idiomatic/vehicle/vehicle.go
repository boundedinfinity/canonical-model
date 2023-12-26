package vehicle

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/color"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Vehicle struct {
	Id             id.Id                       `json:"id,omitempty"`
	Vin            VehicleIdentificationNumber `json:"vin,omitempty"`
	Make           string                      `json:"make,omitempty"`
	Model          string                      `json:"model,omitempty"`
	BodyStyle      BodyStyle                   `json:"body-style,omitempty"`
	BodySize       BodyStyle                   `json:"body-size,omitempty"`
	FuelEfficiency FuelEfficiency              `json:"fuel-efficiency,omitempty"`
	DriveTrain     DriveTrain                  `json:"drive-train,omitempty"`
	FeulType       FeulType                    `json:"fuel-type,omitempty"`
	EngineType     EngineType                  `json:"engine-type,omitempty"`
	Transmission   TransmissionType            `json:"transmission,omitempty"`
	YearModel      rfc3339date.Rfc3339Date     `json:"year-model,omitempty"`
	PurchaseDate   rfc3339date.Rfc3339Date     `json:"purchase-date,omitempty"`
	Features       []Feature                   `json:"features,omitempty"`
	InteriorColor  color.Color                 `json:"interior-color,omitempty"`
	ExteriorColor  color.Color                 `json:"exterior-color,omitempty"`
	Seating        int                         `json:"seating,omitempty"`
}

type FuelEfficiency struct {
	City    int `json:"city,omitempty"`
	Highway int `json:"highway,omitempty"`
}
