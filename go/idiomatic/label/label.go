package label

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewLabel(name, description string) Label {
	m := Label{
		Name:        name,
		Description: description,
	}

	return m
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Label struct {
	Id           ider.Id `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Abbreviation string  `json:"abbreviation,omitempty"`
}
