package payment

type Mail struct {
}

var _ Payment = &Mail{}

func (_ Mail) Kind() Kind {
	return Kinds.Mail
}
