package person

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/certification"
	"github.com/boundedinfinity/rfc3339date"
)

type Dates struct {
	BirthDate      rfc3339date.Rfc3339Date `json:"birth-date,omitempty"`
	DeathDate      rfc3339date.Rfc3339Date `json:"death-date,omitempty"`
	MarriageDates  []MarriageDates         `json:"marriage-dates,omitempty"`
	EducationDates []EducationDates        `json:"dducation-dates,omitempty"`
}

type MarriageDates struct {
	Person PersonModel             `json:"person,omitempty"`
	Start  rfc3339date.Rfc3339Date `json:"start,omitempty"`
	End    rfc3339date.Rfc3339Date `json:"end,omitempty"`
}

type EducationDates struct {
	Institution certification.EducationalInstitution `json:"institution,omitempty"`
	Start       rfc3339date.Rfc3339Date              `json:"start,omitempty"`
	Graduation  rfc3339date.Rfc3339Date              `json:"graduation,omitempty"`
}
