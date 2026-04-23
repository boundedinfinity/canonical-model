package loan

import "errors"

type Rate struct {
	Amount float64
	Period Period
}

var (
	errRate = errors.New("invalid financial rate")
)

func (this Rate) Convert(period Period) Rate {
	var rate Rate

	switch period {
	case Periods.Yearly, Periods.Annual:
		switch this.Period {
		case Periods.Yearly, Periods.Annual:
			rate.Amount = this.Amount
			rate.Period = this.Period
		case Periods.Monthly, Periods.Mensual:
			rate.Amount = this.Amount * 12
			rate.Period = Periods.Yearly
		default:
			panic(errRate)
		}
	case Periods.Monthly, Periods.Mensual:
		switch this.Period {
		case Periods.Yearly, Periods.Annual:
			rate.Amount = this.Amount / 12
			rate.Period = Periods.Monthly
		case Periods.Monthly, Periods.Mensual:
			rate.Amount = this.Amount
			rate.Period = this.Period
		default:
			panic(errRate)
		}
	default:
		panic(errRate)
	}

	return rate
}
