package physical_product

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/measurement"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type PhysicalProduct struct {
	Id           id.Id                  `json:"id,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Manufacturer ManufacturerId         `json:"manufacturer,omitempty"`
	Vendor       VendorId               `json:"vendor,omitempty"`
	Dimensions   measurement.Dimensions `json:"dimensions,omitempty"`
}

type ManufacturerId struct {
	Model  id.ModelNumber  `json:"model-number,omitempty"`
	Serial id.SerialNumber `json:"serial-number,omitempty"`
	Part   id.PartNumber   `json:"part-number,omitempty"`
}

type VendorId struct {
	Sku     id.StockKeepingUnit      `json:"sku,omitempty"`
	Gtin    id.GlobalTradeItemNumber `json:"gtin,omitempty"`
	Upc     id.UniveralProductCode   `json:"upc,omitempty"`
	BarCode id.Barcode               `json:"bar-code,omitempty"`
	QrCode  id.QrCode                `json:"qr-code,omitempty"`
}
