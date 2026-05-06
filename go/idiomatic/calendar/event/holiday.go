package event

import (
	"fmt"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &HolidayEvent{}

func Holiday(r time.DateRange) Event {
	return &HolidayEvent{
		Id:    ider.New(),
		Range: r,
	}
}

type HolidayEvent struct {
	Id       ider.Id
	Title    string
	Range    time.DateRange
	Interval time.Interval
}

func (this HolidayEvent) Kind() Kind {
	return Kinds.Holiday
}

func (this HolidayEvent) String() string {
	return fmt.Sprintf("Holiday: %s", this.Title)
}
