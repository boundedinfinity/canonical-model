package real_estate

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/location"
	"github.com/boundedinfinity/schema/idiomatic/mailing_address"
	"github.com/boundedinfinity/schema/idiomatic/website"
)

type Property struct {
	Id         id.Id                          `json:"id,omitempty"`
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
	Id id.Id `json:"id,omitempty"`
}

type Level struct {
	Id    id.Id  `json:"id,omitempty"`
	Name  id.Id  `json:"name,omitempty"`
	Rooms []Room `json:"rooms,omitempty"`
}

type Room struct {
	Id   id.Id `json:"id,omitempty"`
	Name id.Id `json:"name,omitempty"`
}

type HomeOwnersAssociation struct {
	Id   id.Id `json:"id,omitempty"`
	Name id.Id `json:"name,omitempty"`
}
