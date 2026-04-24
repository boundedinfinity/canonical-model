package show

import (
	"image"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/currency"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/entertainment/details"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/entertainment/fiction"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/entertainment/role"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/language"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location/country"
	"github.com/boundedinfinity/rfc3339date"
)

type Episode struct {
	Id              ider.Id                                  `json:"id"`
	Title           string                                   `json:"title"`
	SubTitles       []string                                 `json:"subtitles"`
	Authors         []role.Author                            `json:"authors"`
	Writers         []role.Writer                            `json:"writers"`
	Actor           []role.Actor                             `json:"actors"`
	ShowRunners     []role.ShowRunner                        `json:"show-runners"`
	Directors       []role.Director                          `json:"directors"`
	Studios         []role.Studio                            `json:"studios"`
	Synopsis        details.Synopsis                         `json:"synopsis"`
	Plot            details.Synopsis                         `json:"plot"`
	StoryLine       details.Synopsis                         `json:"story-line"`
	Setting         details.Setting                          `json:"setting"`
	Language        language.Language                        `json:"language"`
	Posters         []image.Image                            `json:"posters"`
	Rating          []details.MotionPictureAssociationRating `json:"rating"`
	Genres          []fiction.Genre                          `json:"genres"`
	Series          fiction.Series                           `json:"series"`
	Universe        fiction.Universe                         `json:"universe"`
	CountryOfOrigin country.Country                          `json:"country-of-origin"`
	FilmingLocation []details.FilmingLocation                `json:"filming-locatinon"`
	Budget          currency.Amount                          `json:"budget"`
	Runtime         rfc3339date.Rfc3339Duration              `json:"runtime"`
	SoundMixes      []details.SoundMix                       `json:"sound-mixes"`
	ColorMix        []details.ColorMix                       `json:"color-mixes"`
	ContentWarnings []details.ContentWarning                 `json:"content-warnings"`
}
