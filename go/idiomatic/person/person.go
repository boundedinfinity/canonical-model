package person

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location/country"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/person/name"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/professional"
	"github.com/boundedinfinity/rfc3339date"
)

type PersonModel struct {
	Id                      ider.Id                       `json:"id"`
	LegalName               name.Name                     `json:"legal-name"`
	Aliases                 []name.Name                   `json:"aliases"`
	ProfessionalCredentials professional.Credentials      `json:"professional-credentials"`
	DateOfBirth             []rfc3339date.Rfc3339DateTime `json:"date-of-birth"`
	DateOfDeath             []rfc3339date.Rfc3339DateTime `json:"date-of-death"`
	CountryOfBirth          country.Country               `json:"country-of-birth"`
}
