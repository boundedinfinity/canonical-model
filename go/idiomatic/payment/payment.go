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
}

var Kinds = kinds{
	Unknown:                 "unknown",
	ApplePay:                "apple-pay",
	CreditCard:              "credit-card",
	DebitCard:               "debit-card",
	GooglePay:               "google-pay",
	Venmo:                   "venmo",
	PayPal:                  "paypal",
	Phone:                   "phone",
	Mail:                    "mail",
	PrivacyCom:              "privacy.com",
	AutomatedClearningHouse: "ach",
	Zelle:                   "zelle",
	WireTransfer:            "wire-transfer",
	BitCoin:                 "bitcoin",
	WesternUnion:            "western-union",
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Payment interface {
	Kind() Kind
}
