package inventory

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

var _ Item = &PaperItem{}

type PaperItem struct {
	Id   ider.Id   `json:"id"`
	Name string    `json:"name"`
	Size PaperSize `json:"size"`
}

func (this PaperItem) GetName() string {
	return this.Name
}

func (_ PaperItem) Kind() ItemKind {
	return ItemKinds.PaperItem
}
