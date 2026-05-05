package audio_book

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/entertainment/paper_book"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/entertainment/role"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/language"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

type Book struct {
	Id          ider.Id           `json:"id"`
	PaperBook   paper_book.Book   `json:"paper-book"`
	Narrators   []role.Narrator   `json:"narrators"`
	Publisher   role.Publisher    `json:"publisher"`
	Edition     int               `json:"edition"`
	PublishDate time.Date         `json:"publish-date"`
	Duration    time.Duration     `json:"duration"`
	Language    language.Language `json:"language"`
}
