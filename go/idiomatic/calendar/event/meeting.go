package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &MeetingEvent{}

func Meeting(r time.DateRange) Event {
	return &MeetingEvent{
		Id:    ider.New(),
		Range: r,
	}
}

type MeetingEvent struct {
	Id          ider.Id
	Range       time.DateRange
	Description string
	Internal    time.Interval
}

func (this MeetingEvent) Kind() Kind {
	return Kinds.Meeting
}

func (this MeetingEvent) String() string {
	return "Meeting Event: " + this.Description + " (" + this.Range.String() + ")"
}
