package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &FreeEvent{}

func Free(r time.DateRange) Event {
	return &FreeEvent{
		Id:    ider.New(),
		Range: r,
	}
}

type FreeEvent struct {
	Id          ider.Id
	Range       time.DateRange
	Description string
	Internal    time.Interval
}

func (this FreeEvent) Kind() Kind {
	return Kinds.Free
}

func (this FreeEvent) String() string {
	return "Free Time: " + this.Range.String()
}
