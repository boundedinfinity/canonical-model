package audio_book

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/entertainment/paper_book"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/entertainment/role"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type Book struct {
	Id        ider.Id         `json:"id"`
	PaperBook paper_book.Book `json:"paper-book"`
	Narrators []role.Narrator `json:"narrators"`
	Publisher role.Publisher  `json:"publisher"`
	// PublishDate time.Date       `json:"publish-date"`
	// Duration    time.Duration   `json:"duration"`
}
