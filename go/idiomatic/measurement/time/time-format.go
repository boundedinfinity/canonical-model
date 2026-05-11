package time

type TimeFormat string

var TimeFormats = timeFormats{
	Iso8601: "measurement.time.time-format.iso8601",
	Rfc3339: "measurement.time.time-format.rfc3339",
	Unix:    "measurement.time.time-format.unix",
}

type timeFormats struct {
	Iso8601 TimeFormat
	Rfc3339 TimeFormat
	Unix    TimeFormat
}
