package loan

type Loan struct {
	Principle   float64
	DownPayment float64
	Term        Term
	Rate        Rate
}
