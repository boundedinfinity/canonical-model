package calendar

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/calendar/event"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/calendar/visibility"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/label"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

type Calendar struct {
	Id          ider.Id               `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Labels      []label.Labels        `json:"labels"`
	Visibility  visibility.Visibility `json:"visibility"`
	Events      []event.Event         `json:"events"`
}

func (this Calendar) AvailableFor(event event.Event) bool {
	return false
}

func (this Calendar) AvailibilityRange(dates time.DateRange, times time.TimeRange) bool {
	return false
}

func (this Calendar) Add(event event.Event) bool {
	var ok bool
	this.Events = append(this.Events, event)
	return ok
}
