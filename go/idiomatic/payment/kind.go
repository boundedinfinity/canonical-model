package payment

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Kind string

type kinds struct {
	Unknown                 Kind
	ApplePay                Kind
	CreditCard              Kind
	DebitCard               Kind
	GooglePay               Kind
	Venmo                   Kind
	PayPal                  Kind
	Phone                   Kind
	Mail                    Kind
	PrivacyCom              Kind
	AutomatedClearningHouse Kind
	Zelle                   Kind
	WireTransfer            Kind
	BitCoin                 Kind
	WesternUnion            Kind
	Cash                    Kind
	Check                   Kind
	CashiersCheck           Kind
}

var Kinds = kinds{
	Unknown:                 "payment.unknown",
	ApplePay:                "payment.apple-pay",
	CreditCard:              "payment.credit-card",
	DebitCard:               "payment.debit-card",
	GooglePay:               "payment.google-pay",
	Venmo:                   "payment.venmo",
	PayPal:                  "payment.paypal",
	Phone:                   "payment.phone",
	Mail:                    "payment.mail",
	PrivacyCom:              "payment.privacy-com",
	AutomatedClearningHouse: "payment.ach",
	Zelle:                   "payment.zelle",
	WireTransfer:            "payment.wire-transfer",
	BitCoin:                 "payment.bitcoin",
	WesternUnion:            "payment.western-union",
	Cash:                    "payment.cash",
	Check:                   "payment.check",
	CashiersCheck:           "payment.cashiers-check",
}
