package financial

import "errors"

type Loan struct {
	Principle   float64
	DownPayment float64
	Term        FinancialTerm
	Rate        FinancialRate
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
// Period
// ////////////////////////////////////////////////////////////////////////////////////////////////

type FinancialPeriod string

var (
	Annual  FinancialPeriod = "annual"
	Yearly  FinancialPeriod = "yearly"
	Mensual FinancialPeriod = "mensual"
	Monthly FinancialPeriod = "monthly"
)

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
	case Yearly, Annual:
		switch this.Period {
		case Yearly, Annual:
			rate.Amount = this.Amount
			rate.Period = this.Period
		case Monthly, Mensual:
			rate.Amount = this.Amount * 12
			rate.Period = Yearly
		default:
			panic(errFinancialRate)
		}
	case Monthly, Mensual:
		switch this.Period {
		case Yearly, Annual:
			rate.Amount = this.Amount / 12
			rate.Period = Monthly
		case Monthly, Mensual:
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
	case Yearly:
		switch this.Period {
		case Yearly:
			payments.Amount = this.Amount
			payments.Period = this.Period
		case Monthly:
			payments.Amount = this.Amount / 12
			payments.Period = Monthly
		default:
			panic(errFinancialRate)
		}
	case Monthly:
		switch this.Period {
		case Yearly:
			payments.Amount = this.Amount * 12
			payments.Period = Yearly
		case Monthly:
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
