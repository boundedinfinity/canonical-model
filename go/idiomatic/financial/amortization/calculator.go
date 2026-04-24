package amortization

import (
	"math"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/financial/loan"
)

type Item struct {
	Month                  int
	Payment                float64
	Balance                float64
	PrinciplePaid          float64
	InterestPaid           float64
	AcculatedPrinciplePaid float64
	AcculatedInterestPaid  float64
}

type Calculator struct {
	Loan               loan.Loan
	RoundUpPayment     bool // Use this is payment is 1 cent to low
	RoundDownPrinciple bool // Use this is principle is 1 cent to high
}

func (this Calculator) Payment() (float64, error) {
	var result float64
	var err error

	rate := this.Loan.Rate.Convert(loan.Periods.Monthly).Amount / 100
	terms := this.Loan.Term.Convert(loan.Periods.Monthly).Amount
	x1 := math.Pow(float64(1)+rate, float64(terms))
	balance := this.Loan.Principle - this.Loan.DownPayment
	result = balance * rate * x1 / (x1 - 1)

	return result, err
}

func (this Calculator) round(v float64) float64 {
	return math.Round(v*100) / 100
}

func (this Calculator) Calculate() ([]Item, error) {
	var items []Item

	// 	payment, err := this.Payment()

	// 	if err != nil {
	// 		return items, err
	// 	}

	// 	ceil := func(v float64) float64 {
	// 		return math.Ceil(v*100) / 100
	// 	}

	// 	floor := func(v float64) float64 {
	// 		return math.Floor(v*100) / 100
	// 	}

	// 	if this.RoundUpPayment {
	// 		payment = ceil(payment)
	// 	}

	// 	for i := 1; i <= terms; i++ {
	// 		interestPaid := balance * rate
	// 		principlePaid := payment - interestPaid

	// 		if this.RoundDownPrinciple {
	// 			principlePaid = floor(principlePaid)
	// 		}

	// 		balance -= principlePaid

	// 		item := Item{
	// 			Month:         i,
	// 			Payment:       round(payment),
	// 			Balance:       round(balance),
	// 			InterestPaid:  round(interestPaid),
	// 			PrinciplePaid: round(principlePaid),
	// 		}

	// 		items = append(items, item)
	// 	}

	return items, nil
}
