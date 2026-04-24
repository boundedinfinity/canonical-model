package details

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location"
)

type FilmingLocation struct {
	Id       ider.Id `json:"id"`
	Location location.Location
}
