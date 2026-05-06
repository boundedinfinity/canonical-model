package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &TodoEvent{}

func Todo(r time.DateRange) Event {
	return &TodoEvent{
		Id:    ider.New(),
		Range: r,
	}
}

type TodoEvent struct {
	Id          ider.Id
	Range       time.DateRange
	Description string
	Internal    time.Interval
}

func (this TodoEvent) Kind() Kind {
	return Kinds.Todo
}

func (this TodoEvent) String() string {
	return "Todo Event: " + this.Description + " (" + this.Range.String() + ")"
}
