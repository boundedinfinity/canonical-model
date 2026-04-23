package naics

// https://www.census.gov/naics/?58967?yearbck=2022

var Sectors = []*Sector{
	{
		Code: "11",
		Name: "Agriculture, Forestry, Fishing and Hunting",
		SubSectors: []*Sector{
			{
				Code: "111",
				Name: "Crop Production",
				SubSectors: []*Sector{
					{
						Code: "1111",
						Name: "Oilseed and Grain Farming",
						SubSectors: []*Sector{
							{
								Code: "11111",
								Name: "Soybean Farming",
								SubSectors: []*Sector{
									{
										Code: "111110",
										Name: "Soybean Farming",
									},
								},
							},
						},
					},
					{
						Code: "11112",
						Name: "Oilseed (except Soybean) Farming",
						SubSectors: []*Sector{
							{
								Code: "111120",
								Name: "Oilseed (except Soybean) Farming",
							},
						},
					},
					{
						Code: "11113",
						Name: "Dry Pea and Bean Farming",
					},
				},
			},
		},
	},
	{
		Code: "21",
		Name: "Mining, Quarrying, and Oil and Gas Extraction",
	},
	{
		Code: "22",
		Name: "Utilities",
	},
	{
		Code: "23",
		Name: "Construction",
	},
	{
		Code: "31",
		Name: "Manufacturing",
	},
	{
		Code: "42",
		Name: "Wholesale Trade",
	},
	{
		Code: "44",
		Name: "Retail Trade",
	},
	{
		Code: "48",
		Name: "Transportation and Warehousing",
	},
	{
		Code: "51",
		Name: "Information",
	},
	{
		Code: "52",
		Name: "Finance and Insurance",
	},
	{
		Code: "53",
		Name: "Real Estate and Rental and Leasing",
	},
	{
		Code: "54",
		Name: "Professional, Scientific, and Technical Services",
	},
	{
		Code: "55",
		Name: "Management of Companies and Enterprises",
	},
	{
		Code: "56",
		Name: "Administrative and Support and Waste Management and Remediation Services",
	},
	{
		Code: "61",
		Name: "Educational Services",
	},
	{
		Code: "62",
		Name: "Health Care and Social Assistance",
	},
	{
		Code: "71",
		Name: "Arts, Entertainment, and Recreation",
	},
	{
		Code: "72",
		Name: "Accommodation and Food Services",
	},
	{
		Code: "81",
		Name: "Other Services (except Public Administration)",
	},
	{
		Code: "92",
		Name: "Public Administration",
	},
}

func walkSector(parent *Sector, sectors ...*Sector) {
	for _, sector := range sectors {
		if parent != nil {
			sector.Parent = parent
		}

		for _, subSector := range sector.SubSectors {
			walkSector(sector, subSector)
		}
	}
}

func main() {
	walkSector(nil, Sectors...)
}
