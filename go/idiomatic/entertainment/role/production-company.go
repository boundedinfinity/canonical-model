package role

import (
	"image"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type ProductionCompany struct {
	Id       ider.Id           `json:"id"`
	Business business.Business `json:"business"`
	Image    []image.Image     `json:"image"`
}
