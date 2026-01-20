package financial

type InterestRateType string
type InterestRateTerm string

type InterestRate struct {
	Percent float32          `json:"percent,omitempty"`
	Type    InterestRateType `json:"type,omitempty"`
}

type AnnualPercentageRate struct {
	Percent float32 `json:"percent,omitempty"`
}
