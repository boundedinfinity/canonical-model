package vehicle

import "github.com/boundedinfinity/schema/idiomatic/id"

type VehicleIdentificationNumber struct {
	Id  id.Id                       `json:"id,omitempty"`
	Wmi WorldManufacturerIdentifier `json:"wmi,omitempty"`
	Vds VehicleDescriptorSection    `json:"vds,omitempty"`
	Vis VehicleIdentifierSection    `json:"vis,omitempty"`
}

// https://en.wikipedia.org/wiki/Vehicle_identification_number#World_manufacturer_identifier

type WorldManufacturerIdentifier struct {
	Id     id.Id     `json:"id,omitempty"`
	Region WmiRegion `json:"region,omitempty"`
	Other  string    `json:"other,omitempty"`
}

type VehicleDescriptorSection struct {
	Id   id.Id  `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
}

type VehicleIdentifierSection struct {
	Id   id.Id  `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
}
