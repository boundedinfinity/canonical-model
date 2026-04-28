package number

import "github.com/boundedinfinity/go-commoner/errorer"

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	ErrNanp = errorer.New("NANP")
	errNanp = errorer.Func(ErrNumber, ErrNanp)
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ParseNanp(number string) (*NorthAmericanNumberingPlan, error) {
	var nanp NorthAmericanNumberingPlan
	var err error

	return &nanp, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ Number = &NorthAmericanNumberingPlan{}

type NorthAmericanNumberingPlan struct {
	CountryCode       string `json:"country-code"`
	NumberingPlanArea string `json:"numbering-plan-area"`
	CentralOfficeCode string `json:"central-office-code"`
	SubscriberNumber  string `json:"subscriber-number"`
}

func (_ NorthAmericanNumberingPlan) GetKind() Kind {
	return Kinds.Nanp
}

func (this NorthAmericanNumberingPlan) Npa() string {
	return this.NumberingPlanArea
}

func (this NorthAmericanNumberingPlan) AreaCode() string {
	return this.NumberingPlanArea
}

func (this NorthAmericanNumberingPlan) ExchangeCode() string {
	return this.CentralOfficeCode
}

func (this NorthAmericanNumberingPlan) Nxx() string {
	return this.CentralOfficeCode
}

func (this NorthAmericanNumberingPlan) LineNumber() string {
	return this.SubscriberNumber
}

func (this NorthAmericanNumberingPlan) Ln() string {
	return this.SubscriberNumber
}

func (this NorthAmericanNumberingPlan) Sn() string {
	return this.SubscriberNumber
}
