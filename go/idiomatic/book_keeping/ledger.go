package book_keeping

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/label"
)

type LedgerModel struct {
	Id          ider.Id            `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Accounts    []AccountModel     `json:"accounts"`
	Labels      []label.LabelModel `json:"labels"`
}
