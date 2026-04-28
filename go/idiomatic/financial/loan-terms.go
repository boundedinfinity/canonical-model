package financial

import "errors"

type Loan struct {
	Principle   float64
	DownPayment float64
	Term        FinancialTerm
	Rate        FinancialRate
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
// Ratio
// ////////////////////////////////////////////////////////////////////////////////////////////////

type Ratio struct {
	Amount float64
	Period FinancialPeriod
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
// Period
// ////////////////////////////////////////////////////////////////////////////////////////////////

type FinancialPeriod string

var FinancialPeriods = financialPeriods{
	Unknown: "financial.period.unknown",
	Annual:  "financial.period.annual",
	Yearly:  "financial.period.yearly",
	Mensual: "financial.period.mensual",
	Monthly: "financial.period.monthly",
}

type financialPeriods struct {
	Unknown FinancialPeriod
	Annual  FinancialPeriod
	Yearly  FinancialPeriod
	Mensual FinancialPeriod
	Monthly FinancialPeriod
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
// Rate
// ////////////////////////////////////////////////////////////////////////////////////////////////

type FinancialRate struct {
	Amount float64
	Period FinancialPeriod
}

var (
	errFinancialRate = errors.New("invalid financial rate")
)

func (this FinancialRate) Convert(period FinancialPeriod) FinancialRate {
	var rate FinancialRate

	switch period {
	case FinancialPeriods.Yearly, FinancialPeriods.Annual:
		switch this.Period {
		case FinancialPeriods.Yearly, FinancialPeriods.Annual:
			rate.Amount = this.Amount
			rate.Period = this.Period
		case FinancialPeriods.Monthly, FinancialPeriods.Mensual:
			rate.Amount = this.Amount * 12
			rate.Period = FinancialPeriods.Yearly
		default:
			panic(errFinancialRate)
		}
	case FinancialPeriods.Monthly, FinancialPeriods.Mensual:
		switch this.Period {
		case FinancialPeriods.Yearly, FinancialPeriods.Annual:
			rate.Amount = this.Amount / 12
			rate.Period = FinancialPeriods.Monthly
		case FinancialPeriods.Monthly, FinancialPeriods.Mensual:
			rate.Amount = this.Amount
			rate.Period = this.Period
		default:
			panic(errFinancialRate)
		}
	default:
		panic(errFinancialRate)
	}

	return rate
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
// Payments
// ////////////////////////////////////////////////////////////////////////////////////////////////

type FinancialTerm struct {
	Amount int
	Period FinancialPeriod
}

func (this FinancialTerm) Convert(period FinancialPeriod) FinancialTerm {
	var payments FinancialTerm

	switch period {
	case FinancialPeriods.Yearly, FinancialPeriods.Annual:
		switch this.Period {
		case FinancialPeriods.Yearly, FinancialPeriods.Annual:
			payments.Amount = this.Amount
			payments.Period = this.Period
		case FinancialPeriods.Monthly, FinancialPeriods.Mensual:
			payments.Amount = this.Amount / 12
			payments.Period = FinancialPeriods.Monthly
		default:
			panic(errFinancialRate)
		}
	case FinancialPeriods.Monthly, FinancialPeriods.Mensual:
		switch this.Period {
		case FinancialPeriods.Yearly, FinancialPeriods.Annual:
			payments.Amount = this.Amount * 12
			payments.Period = FinancialPeriods.Yearly
		case FinancialPeriods.Monthly, FinancialPeriods.Mensual:
			payments.Amount = this.Amount
			payments.Period = this.Period
		default:
			panic(errFinancialRate)
		}
	default:
		panic(errFinancialRate)
	}

	return payments
}

func (this FinancialTerm) Months() int {
	var months int

	switch this.Period {
	case FinancialPeriods.Yearly, FinancialPeriods.Annual:
		months = this.Amount * 12
	case FinancialPeriods.Monthly, FinancialPeriods.Mensual:
		months = this.Amount
	default:
		panic(errFinancialRate)
	}

	return months
}
