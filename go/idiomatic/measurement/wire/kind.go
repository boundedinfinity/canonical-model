package wire

type Kind string

var Kinds = kinds{
	AmericanWireGauge: "measurement.wire.american-wire-gauge",
}

type kinds struct {
	AmericanWireGauge Kind
}
