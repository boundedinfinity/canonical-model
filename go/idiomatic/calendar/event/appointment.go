package event

import (
	"fmt"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
)

var _ Event = &AppointmentEvent{}

func Appointment(r time.DateRange) Event {
	return &AppointmentEvent{
		Id:    ider.New(),
		Range: r,
	}
}

type AppointmentEvent struct {
	Id    ider.Id
	Range time.DateRange
	Title string
}

func (this AppointmentEvent) Kind() Kind {
	return Kinds.Appointment
}

func (this AppointmentEvent) String() string {
	return fmt.Sprintf("Appointment: %s", this.Title)
}
