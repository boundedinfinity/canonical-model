package holiday

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/calendar/event"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

var _ event.Event = HolidayEventModel{}

// HolidayEventModel represents a holiday.  Typically used for things like Christmas or New Years.
type HolidayEventModel struct {
}

// TODO: This could be pulled from a service.

func (this HolidayEventModel) GetId() ider.Id {
	panic("unimplemented")
}

func (_ HolidayEventModel) GetKind() event.Kind {
	return event.Kinds.Holiday
}
