package contact

import (
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/people"
	"github.com/boundedinfinity/schema/idiomatic/phone"
	"github.com/google/uuid"
)

type Contact struct {
	Id            uuid.UUID          `json:"id,omitempty"`
	Person        people.Person      `json:"person,omitempty"`
	Relationships []people.Person    `json:"relationships,omitempty"`
	Telephones    []phone.NanpNumber `json:"telephones,omitempty"`
	Audit         audit.Audit        `json:"audit,omitempty"`
}
