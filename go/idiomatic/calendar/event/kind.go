package event

type Kind string

type kinds struct {
	Unknown     Kind // Unknown events, such as events with unspecified types.
	Holiday     Kind // Holiday events, such as national holidays, religious holidays, etc.
	Log         Kind // Log events, such as work logs, exercise logs, medication logs, etc.
	Meeting     Kind // Meeting events, such as team meetings, client meetings, etc.
	Span        Kind // Span events, such as vacations, conferences, etc.
	Todo        Kind // Todo events, such as tasks, chores, etc.
	Appointment Kind // Appointment events, such as doctor appointments, dentist appointments, etc.
	Maintenance Kind // Maintenance events, such as car maintenance, home maintenance, etc.
	Significant Kind // Significant events, such as birthdays, marriages, property closing dates, etc.
	Free        Kind // Free events, such as free time, open slots, etc.
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
	Significant: "calendar.event.significant",
}
