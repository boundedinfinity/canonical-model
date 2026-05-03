package angle

type Kind string

var Kinds = kinds{
	Degree: "measurement.angle.degree",
	Radian: "measurement.angle.radian",
}

type kinds struct {
	Degree Kind
	Radian Kind
}
