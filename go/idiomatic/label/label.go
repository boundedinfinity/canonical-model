package label

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewLabel(name, description string) LabelModel {
	m := LabelModel{
		Name:        name,
		Description: description,
	}

	return m
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type LabelModel struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
