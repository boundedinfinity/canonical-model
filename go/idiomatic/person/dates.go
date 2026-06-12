package person

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/certification"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

type Dates struct {
	BirthDate      time.Date        `json:"birth-date,omitempty"`
	DeathDate      time.Date        `json:"death-date,omitempty"`
	MarriageDates  []MarriageDates  `json:"marriage-dates,omitempty"`
	EducationDates []EducationDates `json:"education-dates,omitempty"`
}

type MarriageDates struct {
	Person Person         `json:"person,omitempty"`
	When   time.DateRange `json:"when,omitempty"`
}

type EducationDates struct {
	Institution certification.EducationalInstitution `json:"institution,omitempty"`
	When        time.DateRange                       `json:"when,omitempty"`
}
