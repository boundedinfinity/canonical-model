package role

import (
	"image"

	"github.com/boundedinfinity/canonical_model/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/idiomatic/entertainment/fiction"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

type Actor struct {
	Id         ider.Id              `json:"id"`
	Contact    contact.ContactModel `json:"contact"`
	Image      []image.Image        `json:"image"`
	Synopsis   string               `json:"synopsis"`
	Characters []fiction.Character  `json:"characters"`
}
