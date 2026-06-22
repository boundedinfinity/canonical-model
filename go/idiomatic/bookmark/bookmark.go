package bookmark

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/go-commoner/errorer"
)

// Bookmark represents a bookmark item, including its name, URL, and description.
type Bookmark struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Url         string  `json:"url"`
	Description string  `json:"description"`
}

var ErrBookmark = errorer.New("bookmark error")
