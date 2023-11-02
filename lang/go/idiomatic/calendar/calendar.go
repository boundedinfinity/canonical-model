package calendar

import "github.com/boundedinfinity/schema/idiomatic/id"

// https://en.wikipedia.org/wiki/ICalendar

type Calendar struct {
	Id   id.Id         `json:"id,omitempty"`
	Days []CalendarDay `json:"days"`
}
