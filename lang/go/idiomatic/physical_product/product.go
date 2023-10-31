package physical_product

import "github.com/boundedinfinity/go-commoner/idiomatic/measurement"

type PhysicalProduct struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"string,omitempty"`
	Model       ModelNumber            `json:"model,omitempty"`
	Part        PartNumber             `json:"part,omitempty"`
	Sku         StockKeepingUnit       `json:"sku,omitempty"`
	Serial      SerialNumber           `json:"serial,omitempty"`
	Gtin        GlobalTradeItemNumber  `json:"gtin,omitempty"`
	Upc         UniveralProductCode    `json:"upc,omitempty"`
	BarCode     Barcode                `json:"bar-code,omitempty"`
	QrCode      QrCode                 `json:"qr-code,omitempty"`
	Dimensions  measurement.Dimensions `json:"dimensions,omitempty"`
}
