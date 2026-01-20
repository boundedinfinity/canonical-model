package financial_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/financial"
	"github.com/stretchr/testify/assert"
)

func Test_amortization_1(t *testing.T) {
	loan := financial.Loan{
		Principle:   100000,
		DownPayment: 0,
		Rate:        financial.FinancialRate{Amount: 5, Period: financial.Annual},
		Term:        financial.FinancialTerm{Amount: 30, Period: financial.Yearly},
	}

	calculator := financial.AmortizationCalculator{
		Loan: loan,
	}

	// #	Payment		Interest	Principle	Balance
	// 1	$536.82		$416.67		$120.15		$99,879.85
	// 2	$536.82		$416.17		$120.66		$99,759.19
	// 3	$536.82		$415.66		$121.16		$99,638.03
	// 358	$536.82		$6.65		$530.17		$1,066.97
	// 359	$536.82		$4.45		$532.38		$534.59
	// 360	$536.82		$2.23		$534.59		$0

	items, err := calculator.Calculate()
	assert.Equal(t, len(items), 30)
	assert.NoError(t, err)
}

func Test_amortization_2(t *testing.T) {
	loan := financial.Loan{
		Principle:   130000,
		DownPayment: 32500,
		Rate:        financial.FinancialRate{Amount: 4.125, Period: financial.Annual},
		Term:        financial.FinancialTerm{Amount: 30, Period: financial.Yearly},
	}

	calculator := financial.AmortizationCalculator{
		Loan:               loan,
		RoundUpPayment:     false,
		RoundDownPrinciple: true,
	}

	// #	Payment		Interest	Principle	Balance
	// 1	$472.53		$335.16		$137.77		$97,362.63
	// 2	$472.53		$334.68		$137.85		$97,224.78
	// 3	$472.53		$334.21		$138.32		$97,086.64
	// 358	$472.53		$4.85		$467.68		$942.72
	// 359	$472.53		$3.24		$469.29		$473.43
	// 360	$472.53		$1.63		$473.43		$0

	items, err := calculator.Calculate()
	assert.Equal(t, len(items), 30)
	assert.NoError(t, err)
}

func Test_amortization_3(t *testing.T) {
	loan := financial.Loan{
		Principle:   48000,
		DownPayment: 0,
		Rate:        financial.FinancialRate{Amount: 6.125, Period: financial.Annual},
		Term:        financial.FinancialTerm{Amount: 30, Period: financial.Yearly},
	}

	calculator := financial.AmortizationCalculator{
		Loan:           loan,
		RoundUpPayment: true,
	}

	items, err := calculator.Calculate()
	assert.Equal(t, len(items), 30)
	assert.NoError(t, err)
}
