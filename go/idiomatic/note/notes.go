package note

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type NotesModel struct {
	Id    ider.Id `json:"id"`
	Notes []Note  `json:"notes"`
}
