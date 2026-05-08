package time

type Clock string

var Clocks = clocks{
	Hour24: "measurement.time.clock.24-hour",
	Hour12: "measurement.time.clock.12-hour",
}

type clocks struct {
	Hour24 Clock
	Hour12 Clock
}
