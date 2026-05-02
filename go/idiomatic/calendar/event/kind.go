package event

type Kind string

type kinds struct {
	Unknown     Kind
	Holiday     Kind
	Log         Kind
	Meeting     Kind
	Span        Kind
	Todo        Kind
	Appointment Kind
	Maintenance Kind
	Personal    Kind
	Free        Kind
}

var Kinds = kinds{
	Unknown:     "calendar.event.unknown",
	Holiday:     "calendar.event.holiday",
	Log:         "calendar.event.log",
	Meeting:     "calendar.event.meeting",
	Span:        "calendar.event.span",
	Todo:        "calendar.event.todo",
	Appointment: "calendar.event.appointment",
	Maintenance: "calendar.event.maintenance",
	Free:        "calendar.event.free",
	Personal:    "calendar.event.personal",
}
