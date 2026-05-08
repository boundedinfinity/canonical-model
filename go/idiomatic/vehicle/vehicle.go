package vehicle

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/color"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
	"github.com/boundedinfinity/rfc3339date"
)

type Vehicle struct {
	Id                          ider.Id                     `json:"id,omitempty"`
	Make                        Make                        `json:"make,omitempty"`
	Model                       Model                       `json:"model,omitempty"`
	Vin                         string                      `json:"vin,omitempty"`
	BodyStyle                   BodyStyle                   `json:"body-style,omitempty"`
	BodySize                    BodySize                    `json:"body-size,omitempty"`
	FuelEfficiency              FuelEfficiency              `json:"fuel-efficiency,omitempty"`
	DriveTrain                  DriveTrain                  `json:"drive-train,omitempty"`
	FuelType                    FuelType                    `json:"fuel-type,omitempty"`
	EngineType                  EngineType                  `json:"engine-type,omitempty"`
	Transmission                TransmissionType            `json:"transmission,omitempty"`
	YearModel                   time.Year                   `json:"year-model,omitempty"`
	PurchaseDate                rfc3339date.Rfc3339Date     `json:"purchase-date,omitempty"`
	InteriorColor               color.Color                 `json:"interior-color,omitempty"`
	ExteriorColor               color.Color                 `json:"exterior-color,omitempty"`
	Seating                     int                         `json:"seating,omitempty"`
	VehicleIdentificationNumber VehicleIdentificationNumber `json:"vehicle-identification-number,omitempty"`
	AvailableFeatures           []Feature                   `json:"available-features,omitempty"`
	SelectedFeatures            []Feature                   `json:"selected-features,omitempty"`
}
