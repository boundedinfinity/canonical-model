package details

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/model/location"
)

type FilmingLocation struct {
	Id       ider.Id `json:"id"`
	Location location.Location
}
