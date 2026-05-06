package time

type Interval string

type intervals struct {
	Second     Interval
	Minute     Interval
	Hour       Interval
	Day        Interval
	Week       Interval
	Month      Interval
	Year       Interval
	Decade     Interval
	Century    Interval
	Millennium Interval
}

var Intervals = intervals{
	Second:     "measurement.time.interval.second",
	Minute:     "measurement.time.interval.minute",
	Hour:       "measurement.time.interval.hour",
	Day:        "measurement.time.interval.day",
	Week:       "measurement.time.interval.week",
	Month:      "measurement.time.interval.month",
	Year:       "measurement.time.interval.year",
	Decade:     "measurement.time.interval.decade",
	Century:    "measurement.time.interval.century",
	Millennium: "measurement.time.interval.millennium",
}
