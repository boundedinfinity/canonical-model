package role

import (
	"image"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type Distributor struct {
	Id       ider.Id           `json:"id"`
	Business business.Business `json:"business"`
	Image    []image.Image     `json:"image"`
}
