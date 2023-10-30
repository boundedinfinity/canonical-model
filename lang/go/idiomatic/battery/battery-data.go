package battery

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/math/rational"
	"github.com/boundedinfinity/schema/idiomatic/measurement"
)

var (
	BatteryData = []Battery{
		{ // https://en.wikipedia.org/wiki/AAAA_battery
			Name: "AAAA",
			Dimensions: measurement.Dimensions{
				Diameter: rational.MustString("8.3"),  // mm
				Length:   rational.MustString("42.5"), // mm
			},
		},
	}
)
