package system

type Kind string

type kinds struct {
	Unknown                      Kind
	BritishImperial              Kind
	Metric                       Kind
	SocietyOfAutomotiveEngineers Kind
	UnitedStatesCustomary        Kind
	Unitless                     Kind
}

var Kinds = kinds{
	Unknown:                      "measurement.system.unknown",
	Metric:                       "measurement.system.metric",
	BritishImperial:              "measurement.system.british-imperial",
	Unitless:                     "measurement.system.unitless",
	SocietyOfAutomotiveEngineers: "measurement.system.society-of-automotive-engineers",
	UnitedStatesCustomary:        "measurement.system.united-states-customary",
}
