package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/calendar/visibility"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/label"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/note"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
)

type Event struct {
	Id                  ider.Id               `json:"id"`
	Kind                Kind                  `json:"kind"`
	Title               string                `json:"title"`
	Description         string                `json:"description"`
	Notes               []note.Note           `json:"notes"`
	Mandatory           []person.Person       `json:"mandatory"`
	Optional            []person.Person       `json:"optional"`
	Awareness           []person.Person       `json:"awareness"`
	DateRange           time.DateRange        `json:"date-range"`
	TimeRange           time.TimeRange        `json:"time-range"`
	Labels              label.Labels          `json:"labels"`
	Visibility          visibility.Visibility `json:"visibility"`
	PreparationDuration time.Duration         `json:"preparation-duration"`
	PostDuration        time.Duration         `json:"post-duration"`
	Location            location.Location     `json:"location"`
}
