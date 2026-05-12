package account

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// var _ modeller.Modeller = &AccountModel{}

type Account struct {
	Id     ider.Id `json:"id"`
	Name   string  `json:"name"`
	Number string  `json:"number"`
	// Insitution institution.InstitutionModel `json:"insitution"`
}

func (this *Account) GetId() ider.Id {
	return this.Id
}

func (this *Account) GetName() string {
	return this.Name
}

func (this Account) Last4Digits() string {
	return stringer.TruncateEnd(this.Number, 4, "")
}
