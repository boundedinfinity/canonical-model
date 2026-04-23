package payment

type ApplePay struct {
}

var _ Payment = &ApplePay{}

func (_ ApplePay) Kind() Kind {
	return Kinds.ApplePay
}
