package account

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

// var _ modeller.Modeller = &AccountModel{}

type AccountModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
	// Number     string                       `json:"number"`
	// Insitution institution.InstitutionModel `json:"insitution"`
}

func (this *AccountModel) GetId() ider.Id {
	return this.Id
}

func (this *AccountModel) GetName() string {
	return this.Name
}
