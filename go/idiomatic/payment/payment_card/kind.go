package payment_card

type Kind string

type kinds struct {
	Unknown         Kind
	Debit           Kind
	Credit          Kind
	DigitalPlatform Kind
}

var Kinds = kinds{
	Unknown:         "banking-payment-card.unknown",
	Debit:           "banking-payment-card.debit",
	Credit:          "banking-payment-card.credit",
	DigitalPlatform: "banking-payment-card.digital-platform",
}
