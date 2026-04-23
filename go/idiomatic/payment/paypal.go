package payment

type PayPal struct {
}

var _ Payment = &PayPal{}

func (_ PayPal) Kind() Kind {
	return Kinds.PayPal
}
