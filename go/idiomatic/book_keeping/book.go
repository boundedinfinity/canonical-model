package book_keeping

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type BookModel struct {
	Id          ider.Id       `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Ledgers     []LedgerModel `json:"ledgers"`
}
