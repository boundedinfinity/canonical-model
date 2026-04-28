package physical_product

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

// https://en.wikipedia.org/wiki/Display_resolution

type ScreenDimensions struct {
	Id     ider.Id `json:"id,omitempty"`
	Width  int     `json:"legal-name,omitempty"`
	Height int     `json:"trade-name,omitempty"`
}
