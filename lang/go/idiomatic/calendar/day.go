package calendar

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type CalendarDay struct {
	Id     id.Id                   `json:"id,omitempty"`
	Date   rfc3339date.Rfc3339Date `json:"date"`
	Events []CalendarEvent         `json:"events"`
}
