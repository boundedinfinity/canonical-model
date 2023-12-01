package contact

// https://en.wikipedia.org/wiki/VCard

import (
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/people"
	"github.com/boundedinfinity/schema/idiomatic/phone"
)

type Contact struct {
	Id            id.Id              `json:"id,omitempty"`
	Person        people.Person      `json:"person,omitempty"`
	Relationships []RelationShip     `json:"relationships,omitempty"`
	Telephones    []phone.NanpNumber `json:"telephones,omitempty"`
}
