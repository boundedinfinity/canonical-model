package label

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewLabel(name, description, abbreviation string) Label {
	m := Label{
		Id:           ider.New(),
		Name:         name,
		Description:  optioner.OfZero(description),
		Abbreviation: optioner.OfZero(abbreviation),
	}

	return m
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Label struct {
	Id           ider.Id                 `json:"id"`
	Name         string                  `json:"name"`
	Description  optioner.Option[string] `json:"description"`
	Abbreviation optioner.Option[string] `json:"abbreviation,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BuildLabel(name string) *builder {
	return &builder{
		label: Label{
			Id:   ider.New(),
			Name: name,
		},
	}
}

type builder struct {
	label Label
}

func (this builder) T() Label {
	return this.label
}

func (this *builder) Description(s string) *builder {
	this.label.Description = optioner.OfZero(s)
	return this
}

func (this *builder) Abbreviation(s string) *builder {
	this.label.Abbreviation = optioner.OfZero(s)
	return this
}
