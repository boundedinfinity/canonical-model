package vehicle

type BodySize string

var BodySizes = bodySizes{
	Unknown:  "vehical.body-size.unknown",
	Compact:  "vehical.body-size.compact",
	Midsize:  "vehical.body-size.midsize",
	Fullsize: "vehical.body-size.fullsize",
}

type bodySizes struct {
	Unknown  BodySize
	Compact  BodySize
	Midsize  BodySize
	Fullsize BodySize
}
