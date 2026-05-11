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
	Unknown      CalendarSystem
	Gregorian    CalendarSystem
	Julian       CalendarSystem
	Islamic      CalendarSystem
	Hebrew       CalendarSystem
	Coptic       CalendarSystem
	SolarHijri   CalendarSystem
	VikramSamvat CalendarSystem
	Bengali      CalendarSystem
	UnixEpoch    CalendarSystem
}

var CalendarSystems = calendarSystems{
	Unknown:      "measurement.time.calendar-system.unknown",
	Gregorian:    "measurement.time.calendar-system.gregorian",
	Julian:       "measurement.time.calendar-system.julian",
	Islamic:      "measurement.time.calendar-system.islamic",
	Hebrew:       "measurement.time.calendar-system.hebrew",
	Coptic:       "measurement.time.calendar-system.coptic",
	SolarHijri:   "measurement.time.calendar-system.solar-hijri",
	VikramSamvat: "measurement.time.calendar-system.vikram-samvat",
	Bengali:      "measurement.time.calendar-system.bengali",
}
