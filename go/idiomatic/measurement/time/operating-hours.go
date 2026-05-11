package time

type OperatingHours struct {
	DayOfWeek DayOfWeek   `json:"day-of-week"`
	Hours     TimeRange   `json:"hours"`
	Timezone  TimeZone    `json:"timezone"`
	Clock     ClockFormat `json:"clock,omitempty"`
}
