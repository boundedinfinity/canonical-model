package role

import (
	"image"

	"github.com/boundedinfinity/canonical_model/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

type ShowRunner struct {
	Id      ider.Id              `json:"id"`
	Contact contact.ContactModel `json:"contact"`
	Image   []image.Image        `json:"image"`
}
