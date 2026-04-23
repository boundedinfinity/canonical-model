package book_keeping

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

type ClientModel struct {
	Id          ider.Id                `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Contacts    contact.ContactModel   `json:"contacts"`
	Business    business.BusinessModel `json:"business"`
}
