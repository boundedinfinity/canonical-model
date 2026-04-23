package payment

type GooglePay struct {
}

var _ Payment = &GooglePay{}

func (_ GooglePay) Kind() Kind {
	return Kinds.GooglePay
}
