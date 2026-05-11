package time

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// Month represents a month of the year, such as "January", "February", etc. It can be represented as an interface to allow for different implementations (e.g., using an enum, a string, or a custom struct).
type Month string

func (this Month) String() string {
	return string(this)
}

func (this Month) Number() int {
	switch this {
	case Months.January:
		return 1
	case Months.February:
		return 2
	case Months.March:
		return 3
	case Months.April:
		return 4
	case Months.May:
		return 5
	case Months.June:
		return 6
	case Months.July:
		return 7
	case Months.August:
		return 8
	case Months.September:
		return 9
	case Months.October:
		return 10
	case Months.November:
		return 11
	case Months.December:
		return 12
	default:
		panic("invalid month")
	}
}

var Months = months{
	January:   "measurement.time.month.january",
	February:  "measurement.time.month.february",
	March:     "measurement.time.month.march",
	April:     "measurement.time.month.april",
	May:       "measurement.time.month.may",
	June:      "measurement.time.month.june",
	July:      "measurement.time.month.july",
	August:    "measurement.time.month.august",
	September: "measurement.time.month.september",
	October:   "measurement.time.month.october",
	November:  "measurement.time.month.november",
	December:  "measurement.time.month.december",
}

type months struct {
	January   Month
	February  Month
	March     Month
	April     Month
	May       Month
	June      Month
	July      Month
	August    Month
	September Month
	October   Month
	November  Month
	December  Month
}

var (
	ErrMonth   = errorer.New("month")
	errMonthFn = errorer.Func(ErrMonth)
)

func MonthErr(month int) (Month, error) {
	if month < 1 || month > 12 {
		return "", errMonthFn("month must be between 1 and 12")
	}

	switch month {
	case 1:
		return Months.January, nil
	case 2:
		return Months.February, nil
	case 3:
		return Months.March, nil
	case 4:
		return Months.April, nil
	case 5:
		return Months.May, nil
	case 6:
		return Months.June, nil
	case 7:
		return Months.July, nil
	case 8:
		return Months.August, nil
	case 9:
		return Months.September, nil
	case 10:
		return Months.October, nil
	case 11:
		return Months.November, nil
	case 12:
		return Months.December, nil
	default:
		return "", errMonthFn("invalid month: %d", month)
	}
}

func MonthParseErr(month string) (Month, error) {
	switch stringer.ToLower(month) {
	case "january", "jan", "1", "01":
		return Months.January, nil
	case "february", "feb", "2", "02":
		return Months.February, nil
	case "march", "mar", "3", "03":
		return Months.March, nil
	case "april", "apr", "4", "04":
		return Months.April, nil
	case "may", "5", "05":
		return Months.May, nil
	case "june", "jun", "6", "06":
		return Months.June, nil
	case "july", "jul", "7", "07":
		return Months.July, nil
	case "august", "aug", "8", "08":
		return Months.August, nil
	case "september", "sep", "sept", "9", "09":
		return Months.September, nil
	case "october", "oct", "10":
		return Months.October, nil
	case "november", "nov", "11":
		return Months.November, nil
	case "december", "dec", "12":
		return Months.December, nil
	default:
		return "", errMonthFn("invalid month: %s", month)
	}
}
