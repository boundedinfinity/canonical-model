package location

import (
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type GeoLocation struct {
	Id        id.Id   `json:"id,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
