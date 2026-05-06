package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &SpanEvent{}

func Span(r time.DateRange) Event {
	return &SpanEvent{
		Id:    ider.New(),
		Range: r,
	}
}

type SpanEvent struct {
	Id          ider.Id
	Range       time.DateRange
	Description string
	Internal    time.Interval
}

func (this SpanEvent) Kind() Kind {
	return Kinds.Span
}

func (this SpanEvent) String() string {
	return "Span Event: " + this.Description + " (" + this.Range.String() + ")"
}
