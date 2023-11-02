package calendar

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type CalendarEvent struct {
	Id         id.Id                       `json:"id,omitempty"`
	Organizer  contact.Contact             `json:"organizer"`
	EventStart rfc3339date.Rfc3339DateTime `json:"event-start"`
	EventEnd   rfc3339date.Rfc3339DateTime `json:"event-end"`
	Summary    string                      `json:"summary"`
}
