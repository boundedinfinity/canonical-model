package book_keeping

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/label"
)

type AccountModel struct {
	Id           ider.Id            `json:"id"`
	Name         string             `json:"name"`
	Kind         Kind               `json:"kind"`
	Description  string             `json:"description"`
	Transactions []TransactionModel `json:"transactions"`
	Parent       *AccountModel      `json:"parent"`
	Children     []*AccountModel    `json:"children"`
	Labels       []label.Labels     `json:"labels"`
	Vendor       *VendorModel       `json:"vendor"`
	Client       *ClientModel       `json:"client"`
}
