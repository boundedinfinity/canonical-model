package global_positioning_system

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type GeoLocation struct {
	Id        ider.Id `json:"id,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
