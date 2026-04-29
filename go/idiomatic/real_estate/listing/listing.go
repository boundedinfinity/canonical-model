package listing

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/real_estate/property"
)

type Listing struct {
	Id       ider.Id           `json:"id"`
	Property property.Property `json:"property,omitempty"`
	Price    float64           `json:"price"`
	Images   []string          `json:"images,omitempty"`
}
