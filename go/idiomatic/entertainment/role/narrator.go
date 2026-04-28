package role

import (
	"image"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type Narrator struct {
	Id       ider.Id         `json:"id"`
	Contact  contact.Contact `json:"contact"`
	Image    []image.Image   `json:"image"`
	Synopsis string          `json:"synopsis"`
}
