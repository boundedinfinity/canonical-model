package time

// DayOfYear represents a day of the year, ranging from 1 to 365 (or 366 in a leap year).
type DayOfYear struct {
	Value  int            `json:"value,omitempty"`
	System CalendarSystem `json:"system,omitempty"`
}

// DayOfMonth represents a day of the month, ranging from 1 to 31 depending on the month.
type DayOfMonth struct {
	Value  int            `json:"value,omitempty"`
	System CalendarSystem `json:"system,omitempty"`
}
