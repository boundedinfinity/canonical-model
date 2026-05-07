package humidity

// https://en.wikipedia.org/wiki/Humidity
// Absolute humidity is the mass of water vapor per volume of air (in grams per cubic meter)
// Relative humidity, often expressed as a percentage, indicates a present state of absolute humidity relative to a maximum humidity given the same temperature.
// Specific humidity is the ratio of water vapor mass to total moist air parcel mass.

type Kind string

var Kinds = kinds{
	Absolute: "measurement.humidity.absolute",
	Relative: "measurement.humidity.relative",
	Specific: "measurement.humidity.specific",
}

type kinds struct {
	Absolute Kind
	Relative Kind
	Specific Kind
}
