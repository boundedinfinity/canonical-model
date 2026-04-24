package role

import (
	"image"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type Studio struct {
	Id       ider.Id                `json:"id"`
	Business business.BusinessModel `json:"business"`
	Image    []image.Image          `json:"image"`
}
