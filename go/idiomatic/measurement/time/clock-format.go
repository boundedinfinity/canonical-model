package time

type ClockFormat string

var ClockFormats = clockFormats{
	Hour24: "measurement.time.clock-format.24-hour",
	Hour12: "measurement.time.clock-format.12-hour",
}

type clockFormats struct {
	Hour24 ClockFormat
	Hour12 ClockFormat
}
