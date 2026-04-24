package iso_7812

type MajorIndustryIdentifier struct {
	Category string `json:"category"`
	Digit    string `json:"digit"`
}

var MajorIndustryIdentifiers = majorIndustryIdentifiers{}

type majorIndustryIdentifiers struct {
	all []MajorIndustryIdentifier
}

func (this majorIndustryIdentifiers) Find(code string) (MajorIndustryIdentifier, bool) {
	var ok bool
	var result MajorIndustryIdentifier

	return result, ok
}
