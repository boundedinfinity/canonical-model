package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &SignificantEvent{}

func Significant(r time.DateRange) Event {
	return &SignificantEvent{
		Id:    ider.New(),
		Range: r,
	}
}

type SignificantEvent struct {
	Id          ider.Id
	Range       time.DateRange
	Description string
	Internal    time.Interval
}

func (this SignificantEvent) Kind() Kind {
	return Kinds.Significant
}

func (this SignificantEvent) String() string {
	return "Significant Event: " + this.Description + " (" + this.Range.String() + ")"
}
