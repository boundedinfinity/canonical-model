package vehicle

type EngineType string

var EngineTypes = engineTypes{
	FourCylinder:  "vehical.engine-type.4-cylinder",
	SixCylinder:   "vehical.engine-type.6-cylinder",
	EightCylinder: "vehical.engine-type.8-cylinder",
	OtherCylinder: "vehical.engine-type.other-cylinder",
	Electric:      "vehical.engine-type.electric",
}

type engineTypes struct {
	FourCylinder  EngineType
	SixCylinder   EngineType
	EightCylinder EngineType
	OtherCylinder EngineType
	Electric      EngineType
}
