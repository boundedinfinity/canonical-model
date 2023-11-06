package vehicle

// https://en.wikipedia.org/wiki/Vehicle_identification_number#World_manufacturer_identifier

var (
	codeMap = map[string]WmiRegion{
		"AA": WmiRegions.SouthAfrica,
		"AB": WmiRegions.SouthAfrica,
		"AC": WmiRegions.SouthAfrica,
		"AD": WmiRegions.SouthAfrica,
		"AE": WmiRegions.SouthAfrica,
		"AF": WmiRegions.SouthAfrica,
		"AG": WmiRegions.SouthAfrica,
		"AH": WmiRegions.SouthAfrica,
		"AJ": WmiRegions.CoteDIvoire,
		"AK": WmiRegions.CoteDIvoire,
		"AL": WmiRegions.CoteDIvoire,
		"AM": WmiRegions.CoteDIvoire,
		"AN": WmiRegions.CoteDIvoire,
	}
)
