package appointment

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/calendar/event"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

var _ event.Event = AppointmentEventModel{}

// AppointmentEventModel represents an appointment.  Typically used for things like a doctors appointment or oil change.
type AppointmentEventModel struct {
}

func (this AppointmentEventModel) GetId() ider.Id {
	panic("unimplemented")
}

func (_ AppointmentEventModel) GetKind() event.Kind {
	return event.Kinds.Appointment
}
