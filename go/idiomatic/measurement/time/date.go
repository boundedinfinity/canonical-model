package time

import gtime "time"

type Date struct {
	Year   Year
	Month  Month
	Day    DayOfMonth
	Format TimeFormat
}

func (this Date) String() string {
	return ""

}

func (this Date) Native() gtime.Time {
	return gtime.Date(int(this.Year), gtime.Month(this.Month.native), int(this.Day), 0, 0, 0, 0, gtime.UTC)
}

func (this Date) IsZero() bool {
	return true
}

func (this Date) Before(date Date) bool {
	return true
}

func (this Date) After(date Date) bool {
	return true
}

func (this Date) Equal(date Date) bool {
	return true
}
