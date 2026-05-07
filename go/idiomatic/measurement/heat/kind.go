package heat

// https://en.wikipedia.org/wiki/Heat

type Kind string

var Kinds = kinds{
	Joule:              "measurement.heat.joule",
	BritishThermalUnit: "measurement.heat.british-thermal-unit",
	Calorie:            "measurement.heat.colorie",
}

type kinds struct {
	Joule              Kind
	BritishThermalUnit Kind
	Calorie            Kind
}
