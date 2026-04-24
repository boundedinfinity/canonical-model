package show

import (
	"image"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/entertainment/details"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/entertainment/fiction"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/entertainment/role"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/language"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location/country"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/measurement/ratio"
)

type Show struct {
	Id              ider.Id           `json:"id"`
	Title           string            `json:"title"`
	SubTitles       []string          `json:"subtitles"`
	Creators        []role.Creator    `json:"creators"`
	Writers         []role.Producer   `json:"writers"`
	Studios         []role.Studio     `json:"studios"`
	Synopsis        string            `json:"synopsis"`
	Language        language.Language `json:"language"`
	Posters         []image.Image     `json:"posters"`
	Genres          []fiction.Genre   `json:"genres"`
	Series          fiction.Series    `json:"series"`
	Universe        fiction.Universe  `json:"universe"`
	CountryOfOrigin country.Country   `json:"country-of-origin"`
	Seasons         Season            `json:"seasons"`
	Rating          details.TvRating  `json:"tv-rating"`
	AspectRatio     ratio.AspectRatio `json:"aspect-ratio"`
}
