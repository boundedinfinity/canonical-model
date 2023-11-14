package banking

var FederalReserveDistrictInfo = federalReserveDistrictInfo{}

type federalReserveDistrictInfo struct {
	number map[int]FederalReserveDistrict
	order  map[string]FederalReserveDistrict
	letter map[string]FederalReserveDistrict
}

// func (t federalReserveDistrictInfo) FromNo(number int) FederalReserveDistrict {

// }

func init() {
	FederalReserveDistrictInfo.number = map[int]FederalReserveDistrict{
		1:  FederalReserveDistricts.Boston,
		2:  FederalReserveDistricts.NewYork,
		3:  FederalReserveDistricts.Philadelphia,
		4:  FederalReserveDistricts.Cleveland,
		5:  FederalReserveDistricts.Richmond,
		6:  FederalReserveDistricts.Atlanta,
		7:  FederalReserveDistricts.Chicago,
		8:  FederalReserveDistricts.StLouis,
		9:  FederalReserveDistricts.Minneapolis,
		10: FederalReserveDistricts.KansasCity,
		11: FederalReserveDistricts.Dallas,
		12: FederalReserveDistricts.SanFrancisco,
	}

	FederalReserveDistrictInfo.order = map[string]FederalReserveDistrict{
		"1st":  FederalReserveDistricts.Boston,
		"2nd":  FederalReserveDistricts.NewYork,
		"3rd":  FederalReserveDistricts.Philadelphia,
		"4th":  FederalReserveDistricts.Cleveland,
		"5th":  FederalReserveDistricts.Richmond,
		"6th":  FederalReserveDistricts.Atlanta,
		"7th":  FederalReserveDistricts.Chicago,
		"8th":  FederalReserveDistricts.StLouis,
		"9th":  FederalReserveDistricts.Minneapolis,
		"10th": FederalReserveDistricts.KansasCity,
		"11th": FederalReserveDistricts.Dallas,
		"12th": FederalReserveDistricts.SanFrancisco,
	}
	FederalReserveDistrictInfo.letter = map[string]FederalReserveDistrict{
		"A": FederalReserveDistricts.Boston,
		"B": FederalReserveDistricts.NewYork,
		"C": FederalReserveDistricts.Philadelphia,
		"D": FederalReserveDistricts.Cleveland,
		"E": FederalReserveDistricts.Richmond,
		"F": FederalReserveDistricts.Atlanta,
		"G": FederalReserveDistricts.Chicago,
		"H": FederalReserveDistricts.StLouis,
		"I": FederalReserveDistricts.Minneapolis,
		"J": FederalReserveDistricts.KansasCity,
		"K": FederalReserveDistricts.Dallas,
		"L": FederalReserveDistricts.SanFrancisco,
	}
}
