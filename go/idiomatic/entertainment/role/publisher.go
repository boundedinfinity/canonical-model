package role

import (
	"image"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type Publisher struct {
	Id       ider.Id           `json:"id"`
	Business business.Business `json:"business"`
	Image    []image.Image     `json:"image"`
}
