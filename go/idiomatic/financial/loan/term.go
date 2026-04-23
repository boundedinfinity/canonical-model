package loan

type Term struct {
	Amount int
	Period Period
}

func (this Term) Convert(period Period) Term {
	var payments Term

	switch period {
	case Periods.Yearly:
		switch this.Period {
		case Periods.Yearly:
			payments.Amount = this.Amount
			payments.Period = this.Period
		case Periods.Monthly:
			payments.Amount = this.Amount / 12
			payments.Period = Periods.Monthly
		default:
			panic(errRate)
		}
	case Periods.Monthly:
		switch this.Period {
		case Periods.Yearly:
			payments.Amount = this.Amount * 12
			payments.Period = Periods.Yearly
		case Periods.Monthly:
			payments.Amount = this.Amount
			payments.Period = this.Period
		default:
			panic(errRate)
		}
	default:
		panic(errRate)
	}

	return payments
}
