package time

import (
	"github.com/boundedinfinity/go-commoner/errorer"
)

// DayOfYear represents a day of the year, ranging from 1 to 365 (or 366 in a leap year).
type DayOfYear struct {
	Value  int            `json:"value,omitempty"`
	System CalendarSystem `json:"system,omitempty"`
}

var (
	ErrDayOfMonth   = errorer.New("day of month")
	errDayOfMonthFn = errorer.Func(ErrDayOfMonth)
)

// DayOfMonth represents a day of the month, ranging from 1 to 31 depending on the month.
type DayOfMonth int

// DayOfMonthErr validates the day of the month based on the given year and month, accounting for leap years and varying month lengths.
// It returns an error if the day is invalid for the specified month and year.
func DayOfMonthErr(year Year, month Month, day int) (DayOfMonth, error) {
	if day < 1 {
		return 0, errDayOfMonthFn("less than 1")
	}

	// Months.January:   31,
	// Months.February:  28, // 29 in a leap year
	// Months.March:     31,
	// Months.April:     30,
	// Months.May:       31,
	// Months.June:      30,
	// Months.July:      31,
	// Months.August:    31,
	// Months.September: 30,
	// Months.October:   31,
	// Months.November:  30,
	// Months.December:  31,

	switch month {
	case Months.January, Months.March, Months.May, Months.July, Months.August, Months.October, Months.December:
		if day > 31 {
			return 0, errDayOfMonthFn("greater than 31 in month of %s", month)
		}
	case Months.April, Months.June, Months.September, Months.November:
		if day > 30 {
			return 0, errDayOfMonthFn("greater than 30 in month of %s", month)
		}
	case Months.February:
		if day > 29 {
			return 0, errDayOfMonthFn("greater than 29 in month of %s", month)
		}
	}

	return DayOfMonth(day), nil
}
