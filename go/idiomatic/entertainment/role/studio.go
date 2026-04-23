package role

import (
	"image"

	"github.com/boundedinfinity/canonical_model/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

type Studio struct {
	Id       ider.Id                `json:"id"`
	Business business.BusinessModel `json:"business"`
	Image    []image.Image          `json:"image"`
}
