package book_keeping

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/label"
)

type VendorModel struct {
	Id          ider.Id                `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Contacts    contact.ContactModel   `json:"contacts"`
	Business    business.BusinessModel `json:"business"`
	Labels      []label.LabelModel     `json:"labels"`
}
