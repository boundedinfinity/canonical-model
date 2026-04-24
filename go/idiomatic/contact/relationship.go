package contact

type RelationShip struct {
	Id      ider.Id          `json:"id,omitempty"`
	Type    RelationshipType `json:"relationship-type,omitempty"`
	Contact RelationshipType `json:"contact,omitempty"`
}

type RelationshipType struct {
	Id   ider.Id `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}
