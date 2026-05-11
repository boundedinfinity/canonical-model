package time

// Age represents time ages like "iron age", "bronze age", "stone age", etc. It includes a name and a date range.
type Age struct {
	Name  string    `json:"name,omitempty"`
	Range YearRange `json:"range,omitempty"`
}

var Ages = ages{
	PreHistoric:        Age{Name: "Pre Historic", Range: YearRange{End: Year(-2500000)}},
	Paleolithic:        Age{Name: "Paleolithic", Range: YearRange{Start: Year(-2500000), End: Year(-10000)}},
	Mesolithic:         Age{Name: "Mesolithic", Range: YearRange{Start: Year(-10000), End: Year(-8000)}},
	Neolithic:          Age{Name: "Neolithic", Range: YearRange{Start: Year(-8000), End: Year(-3000)}},
	Bronze:             Age{Name: "Bronze Age", Range: YearRange{Start: Year(-3000), End: Year(-1200)}},
	Iron:               Age{Name: "Iron Age", Range: YearRange{Start: Year(-1200), End: Year(500)}},
	ClassicalAntiquity: Age{Name: "Classical Antiquity", Range: YearRange{Start: Year(-800), End: Year(500)}},
	MiddleAges:         Age{Name: "Middle Ages", Range: YearRange{Start: Year(500), End: Year(1500)}},
	Renaissance:        Age{Name: "Renaissance", Range: YearRange{Start: Year(1300), End: Year(1600)}},
	Modern:             Age{Name: "Modern Era", Range: YearRange{Start: Year(1500), End: Year(1900)}},
	Contemporary:       Age{Name: "Contemporary Era", Range: YearRange{Start: Year(1900), End: Year(3000)}},
}

type ages struct {
	PreHistoric        Age
	Paleolithic        Age
	Mesolithic         Age
	Neolithic          Age
	Bronze             Age
	Iron               Age
	ClassicalAntiquity Age
	MiddleAges         Age
	Renaissance        Age
	Modern             Age
	Contemporary       Age
}
