package contact

import "github.com/boundedinfinity/schema/idiomatic/id"

type RelationShip struct {
	Id      id.Id            `json:"id,omitempty"`
	Type    RelationshipType `json:"relationship-type,omitempty"`
	Contact RelationshipType `json:"contact,omitempty"`
}

type RelationshipType struct {
	Id   id.Id  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
