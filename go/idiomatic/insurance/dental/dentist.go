package dental

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/person"
)

type Dentist struct {
	Id        ider.Id            `json:"id,omitempty"`
	Dentist   person.PersonModel `json:"dentist,omitempty"`
	DentistId string             `json:"dentist-id,omitempty"`
}
