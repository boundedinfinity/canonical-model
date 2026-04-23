package time

// https://en.wikipedia.org/wiki/Gregorian_calendar
// https://en.wikipedia.org/wiki/Julian_calendar
// https://en.wikipedia.org/wiki/Hebrew_calendar
// https://en.wikipedia.org/wiki/Coptic_calendar
// https://en.wikipedia.org/wiki/Solar_Hijri_calendar
// https://en.wikipedia.org/wiki/Bengali_calendar
// https://en.wikipedia.org/wiki/Islamic_calendar

type CalendarSystem string

type calendarSystems struct {
	Unknown    CalendarSystem
	Gregorian  CalendarSystem
	Julian     CalendarSystem
	Islamic    CalendarSystem
	Hebrew     CalendarSystem
	Coptic     CalendarSystem
	SolarHijri CalendarSystem
	Bengali    CalendarSystem
}

var CalendarSystems = calendarSystems{
	Unknown:    "time.calendar-system.unknown",
	Gregorian:  "time.calendar-system.gregorian",
	Julian:     "time.calendar-system.julian",
	Islamic:    "time.calendar-system.islamic",
	Hebrew:     "time.calendar-system.hebrew",
	Coptic:     "time.calendar-system.coptic",
	SolarHijri: "time.calendar-system.solar_hijri",
	Bengali:    "time.calendar-system.bengali",
}
