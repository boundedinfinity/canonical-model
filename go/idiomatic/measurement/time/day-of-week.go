package time

type DayOfWeek string

type daysOfWeek struct {
	Unknown   DayOfWeek
	Monday    DayOfWeek
	Tuesday   DayOfWeek
	Wednesday DayOfWeek
	Thursday  DayOfWeek
	Friday    DayOfWeek
	Saturday  DayOfWeek
	Sunday    DayOfWeek
}

var DaysOfWeek = daysOfWeek{
	Unknown:   "time.day-of-week.unknown",
	Monday:    "time.day-of-week.monday",
	Tuesday:   "time.day-of-week.tuesday",
	Wednesday: "time.day-of-week.wednesday",
	Thursday:  "time.day-of-week.thursday",
	Friday:    "time.day-of-week.friday",
	Saturday:  "time.day-of-week.saturday",
	Sunday:    "time.day-of-week.sunday",
}
