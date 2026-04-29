package physical_product

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type PhysicalProduct struct {
	Id           ider.Id        `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Description  string         `json:"description,omitempty"`
	Manufacturer ManufacturerId `json:"manufacturer,omitempty"`
	Vendor       VendorId       `json:"vendor,omitempty"`
	// Dimensions   measurement.Dimensions `json:"dimensions,omitempty"`
}

type ManufacturerId struct {
	Model  string `json:"model-number,omitempty"`
	Serial string `json:"serial-number,omitempty"`
	Part   string `json:"part-number,omitempty"`
}

type VendorId struct {
	Sku     string `json:"sku,omitempty"`
	Gtin    string `json:"gtin,omitempty"`
	Upc     string `json:"upc,omitempty"`
	BarCode string `json:"bar-code,omitempty"`
	QrCode  string `json:"qr-code,omitempty"`
}
