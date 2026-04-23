package resource

type Kind string

type kinds struct {
	Unknown       Kind
	Person        Kind
	Contact       Kind
	Calendar      Kind
	CalendarEvent Kind
}

var Kinds = kinds{
	Unknown:       "digital.access-management.resource.unknown",
	Person:        "digital.access-management.resource.person",
	Contact:       "digital.access-management.resource.contact",
	Calendar:      "digital.access-management.resource.calendar",
	CalendarEvent: "digital.access-management.resource.calendar-event",
}
