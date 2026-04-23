package time

type BusinessHours struct {
	DayOfWeek DayOfWeek `json:"day-of-week"`
	Hours     TimeRange `json:"hours"`
	Timezone  TimeZone  `json:"timezone"`
}
