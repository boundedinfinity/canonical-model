package movie

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/model/location/country"
	"github.com/boundedinfinity/canonical_model/idiomatic/model/specifications/rfc/rfc3339date"
)

type TheatricalRelease struct {
	ReleaseDate rfc3339date.Rfc3339Date `json:"release-date"`
	Country     country.Country         `json:"countries"`
}
