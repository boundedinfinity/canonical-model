package vehicle

import (
	"errors"
)

type VehicleIdentificationNumber struct {
	Wmi WorldManufacturerIdentifier `json:"wmi,omitempty"`
	Vds VehicleDescriptorSection    `json:"vds,omitempty"`
	Vis VehicleIdentifierSection    `json:"vis,omitempty"`
}

// https://en.wikipedia.org/wiki/Vehicle_identification_number#World_manufacturer_identifier

type WorldManufacturerIdentifier struct {
	Region WmiRegion `json:"region,omitempty"`
	Other  string    `json:"other,omitempty"`
}

type VehicleDescriptorSection struct {
	Code string `json:"code,omitempty"`
}

type VehicleIdentifierSection struct {
	Code string `json:"code,omitempty"`
}

var VehicleIdentificationNumbers = vehicleIdentificationNumbers{}

type vehicleIdentificationNumbers struct {
	ErrInvalidVin error
}

func init() {
	VehicleIdentificationNumbers.ErrInvalidVin = errors.New("invalid VIN")
}

func (t vehicleIdentificationNumbers) Parse(input string) (WorldManufacturerIdentifier, error) {
	var vin WorldManufacturerIdentifier

	if len(input) != 17 {
		return vin, t.ErrInvalidVin
	}

	return vin, nil
}

func (t vehicleIdentificationNumbers) MustParse(input string) WorldManufacturerIdentifier {
	vin, err := t.Parse(input)

	if err != nil {
		panic(err)
	}

	return vin
}
