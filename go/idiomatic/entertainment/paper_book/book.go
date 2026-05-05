package paper_book

import (
	"image"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/currency"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/entertainment/fiction"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/entertainment/role"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/language"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

type Book struct {
	Id                              ider.Id                         `json:"id"`
	Title                           string                          `json:"title"`
	SubTitles                       []string                        `json:"subtitles"`
	Authors                         []role.Author                   `json:"authors"`
	Editors                         []role.Editor                   `json:"editors"`
	InternationalStandardBookNumber InternationalStandardBookNumber `json:"international-standard-book-number"`
	Edition                         int                             `json:"edition"`
	ListPrice                       currency.Amount                 `json:"list-price"`
	Publisher                       role.Publisher                  `json:"publisher"`
	PublishDate                     time.Date                       `json:"publish-date"`
	PageCount                       int                             `json:"page-count"`
	WordCount                       int                             `json:"word-count"`
	Synopsis                        string                          `json:"synopsis"`
	Language                        language.Language               `json:"language"`
	FrontCover                      []image.Image                   `json:"front-cover"`
	BackCover                       []image.Image                   `json:"back-cover"`
	Characters                      []fiction.Character             `json:"characters"`
	Genres                          []fiction.Genre                 `json:"genres"`
	Series                          fiction.Series                  `json:"series"`
	Universe                        fiction.Universe                `json:"universe"`
}

func (this Book) AuthorAgeAtTimeOfPublication() int {
	return 0
}
