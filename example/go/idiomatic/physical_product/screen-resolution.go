package physical_product

import "github.com/boundedinfinity/schema/idiomatic/id"

// https://en.wikipedia.org/wiki/Display_resolution

type ScreenDimensions struct {
	Id     id.Id `json:"id,omitempty"`
	Width  int   `json:"legal-name,omitempty"`
	Height int   `json:"trade-name,omitempty"`
}
