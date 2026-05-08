package vehicle

type BodyStyle string

var BodyStyles = bodyStyles{
	Unknown:             "vehical.body-style.unknown",
	Sedan:               "vehical.body-style.sedan",
	Coupe:               "vehical.body-style.coupe",
	Hatchback:           "vehical.body-style.hatchback",
	SportUtilityVehicle: "vehical.body-style.sport-utility-vehicle",
	Truck:               "vehical.body-style.truck",
	Van:                 "vehical.body-style.van",
	Minivan:             "vehical.body-style.minivan",
	Convertible:         "vehical.body-style.convertible",
	Wagon:               "vehical.body-style.wagon",
	SportWagon:          "vehical.body-style.sport-wagon",
}

type bodyStyles struct {
	Unknown             BodyStyle
	Sedan               BodyStyle
	Coupe               BodyStyle
	Hatchback           BodyStyle
	SportUtilityVehicle BodyStyle
	Truck               BodyStyle
	Van                 BodyStyle
	Minivan             BodyStyle
	Convertible         BodyStyle
	Wagon               BodyStyle
	SportWagon          BodyStyle
}
