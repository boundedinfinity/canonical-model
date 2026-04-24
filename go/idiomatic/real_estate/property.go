package real_estate

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/mailing_address"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/website"
	"github.com/boundedinfinity/rfc3339date"
)

type Property struct {
	Id         ider.Id                        `json:"id,omitempty"`
	Address    mailing_address.MailingAddress `json:"address,omitempty"`
	Parcel     location.Parcel                `json:"parcel,omitempty"`
	FloorPlan  FloorPlan                      `json:"floor-plan,omitempty"`
	Hoa        HomeOwnersAssociation          `json:"hoa,omitempty"`
	BuiltIn    rfc3339date.Rfc3339Date        `json:"built-in,omitempty"`
	Levels     []Level                        `json:"levels,omitempty"`
	Parking    Parking                        `json:"parking,omitempty"`
	References website.PortalWebSite          `json:"references,omitempty"`
}

type Parking struct {
	Id ider.Id `json:"id,omitempty"`
}

type Level struct {
	Id    ider.Id `json:"id,omitempty"`
	Name  ider.Id `json:"name,omitempty"`
	Rooms []Room  `json:"rooms,omitempty"`
}

type Room struct {
	Id   ider.Id `json:"id,omitempty"`
	Name ider.Id `json:"name,omitempty"`
}

type HomeOwnersAssociation struct {
	Id   ider.Id `json:"id,omitempty"`
	Name ider.Id `json:"name,omitempty"`
}
