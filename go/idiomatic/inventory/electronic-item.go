package inventory

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

var _ Item = &PaperItem{}

type ElectronicItem struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this ElectronicItem) GetName() string {
	return this.Name
}

func (_ ElectronicItem) Kind() ItemKind {
	return ItemKinds.ElectronicItem
}
