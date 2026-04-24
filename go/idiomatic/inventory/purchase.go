package inventory

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/currency"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location/mailing_address"
)

type Purchase struct {
	Id                 ider.Id                 `json:"id"`
	OrderNumber        string                  `json:"order-number"`
	ConfirmationNumber string                  `json:"confirmation-number"`
	TrackingNumber     string                  `json:"tracking-number"`
	Cart               Cart                    `json:"cart"`
	BillingAddress     mailing_address.Address `json:"billing-address"`
	ShippingAddress    mailing_address.Address `json:"shipping-address"`
	Vendor             business.BusinessModel  `json:"vendor"`
}

type Cart struct {
	Items    []CartItem      `json:"items"`
	Shipping currency.Amount `json:"shipping"`
	Tax      currency.Amount `json:"tax"`
}

type CartItem struct {
	PhysicalItem PhysicalItem    `json:"physical-item"`
	Quantity     int             `json:"quantity"`
	Price        currency.Amount `json:"price"`
}
