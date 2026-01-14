package banking

type FederalReserveDistrict struct {
	Name   string
	Number int
	Letter string
}

var FederalReserveDistricts = federalReserveDistricts{}

type federalReserveDistricts struct {
	all    []FederalReserveDistrict
	number map[int]FederalReserveDistrict
	letter map[string]FederalReserveDistrict

	Boston       FederalReserveDistrict
	NewYork      FederalReserveDistrict
	Philadelphia FederalReserveDistrict
	Cleveland    FederalReserveDistrict
	Richmond     FederalReserveDistrict
	Atlanta      FederalReserveDistrict
	Chicago      FederalReserveDistrict
	StLouis      FederalReserveDistrict
	Minneapolis  FederalReserveDistrict
	KansasCity   FederalReserveDistrict
	Dallas       FederalReserveDistrict
	SanFrancisco FederalReserveDistrict
}

func (this federalReserveDistricts) All() []FederalReserveDistrict {
	return this.all
}

func init() {
	FederalReserveDistricts.Boston = FederalReserveDistrict{
		Name:   "Federal Reserve Bank of Boston",
		Number: 1,
		Letter: "A",
	}

	FederalReserveDistricts.Boston = FederalReserveDistrict{
		Name:   "Federal Reserve Bank of New York",
		Number: 2,
		Letter: "B",
	}

	FederalReserveDistricts.Boston = FederalReserveDistrict{
		Name:   "Federal Reserve Bank of Philadelphia",
		Number: 3,
		Letter: "C",
	}

	FederalReserveDistricts.Boston = FederalReserveDistrict{
		Name:   "Federal Reserve Bank of Clevland",
		Number: 4,
		Letter: "D",
	}

	FederalReserveDistricts.all = []FederalReserveDistrict{
		FederalReserveDistricts.Boston,
		FederalReserveDistricts.NewYork,
		FederalReserveDistricts.Philadelphia,
		FederalReserveDistricts.Cleveland,
		FederalReserveDistricts.Richmond,
		FederalReserveDistricts.Atlanta,
		FederalReserveDistricts.Chicago,
		FederalReserveDistricts.StLouis,
		FederalReserveDistricts.Minneapolis,
		FederalReserveDistricts.KansasCity,
		FederalReserveDistricts.Dallas,
		FederalReserveDistricts.SanFrancisco,
	}

	for _, district := range FederalReserveDistricts.all {
		FederalReserveDistricts.number[district.Number] = district
	}

	for _, district := range FederalReserveDistricts.all {
		FederalReserveDistricts.letter[district.Letter] = district
	}
}
