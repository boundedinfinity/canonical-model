package vehicle

type FuelType string

var FuelTypes = fuelTypes{
	Unknown:      "vehical.fuel-type.unknown",
	Diesel:       "vehical.fuel-type.diesel",
	Electric:     "vehical.fuel-type.electric",
	Gas:          "vehical.fuel-type.gas",
	Hybrid:       "vehical.fuel-type.hybrid",
	Other:        "vehical.fuel-type.other",
	PlugInHybrid: "vehical.fuel-type.plug-in-hybrid",
}

type fuelTypes struct {
	Unknown      FuelType
	Diesel       FuelType
	Electric     FuelType
	Gas          FuelType
	Hybrid       FuelType
	Other        FuelType
	PlugInHybrid FuelType
}
