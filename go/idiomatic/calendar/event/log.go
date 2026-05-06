package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &LogEvent{}

func LoggerNow(title string) func() Event {
	return func() Event {
		return LogNow(title)
	}
}

func LogNow(title string) Event {
	return &LogEvent{
		Id:    ider.New(),
		Range: time.DateRange{}, // Assuming LogNow uses the current date range
		Title: title,
	}
}

func Log(title string, r time.DateRange) Event {
	return &LogEvent{
		Id:    ider.New(),
		Range: r,
		Title: title,
	}
}

type LogEvent struct {
	Id    ider.Id
	Title string
	Range time.DateRange
}

func (this LogEvent) Kind() Kind {
	return Kinds.Log
}

func (this LogEvent) String() string {
	return "Log Event: " + this.Title + " (" + this.Range.String() + ")"
}
