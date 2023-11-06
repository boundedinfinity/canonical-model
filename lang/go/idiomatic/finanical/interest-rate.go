package finanical

type InterestRateTerm string

type InterestRate struct {
	Term    InterestRateTerm `json:"term,omitempty"`
	Percent float32          `json:"percent,omitempty"`
}
