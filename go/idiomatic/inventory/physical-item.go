package inventory

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/idiomatic/digital/bookmark"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ Item = &PhysicalItem{}

type PhysicalItem struct {
	Id                            ider.Id                 `json:"id"`
	Name                          string                  `json:"name"`
	StokeKeepingUnit              StockKeepingUnit        `json:"stoke-keeping-unit"`
	Model                         string                  `json:"model"`
	SerialNumber                  string                  `json:"serial-number"`
	ProduceCode                   string                  `json:"produce-code"`
	ManufactureDate               rfc3339date.Rfc3339Date `json:"manufacture-date"`
	PurchaseDate                  rfc3339date.Rfc3339Date `json:"purchase-date"`
	Vendor                        business.BusinessModel  `json:"vendor"`
	ProducatPage                  bookmark.Bookmark       `json:"producat-page"`
	SupportPage                   bookmark.Bookmark       `json:"support-page"`
	OriginalEquipmentManufacturer business.BusinessModel  `json:"original-equipment-manufacturer"`
}

func (this PhysicalItem) GetName() string {
	return this.Name
}

func (_ PhysicalItem) Kind() ItemKind {
	return ItemKinds.PhysicalItem
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
