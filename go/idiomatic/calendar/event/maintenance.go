package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &MaintenanceEvent{}

func Maintenance(r time.DateRange) Event {
	return &MaintenanceEvent{
		Id:    ider.New(),
		Range: r,
	}
}

// Maintenance represents a maintenance event.  This is similar to the todo but requires more explaination
// and has a number steps that need to be completed. Think of this like and Standard Operating Procedure (SOP).
type MaintenanceEvent struct {
	Id          ider.Id
	Range       time.DateRange
	Description string
	Internal    time.Interval
}

func (this MaintenanceEvent) Kind() Kind {
	return Kinds.Maintenance
}

func (this MaintenanceEvent) String() string {
	return "Maintenance Event: " + this.Description + " (" + this.Range.String() + ")"
}
