package payment

type Mail struct {
}

var _ Payment = &Mail{}

func (_ Mail) Kind() Kind {
	return Kinds.Mail
}

func (m Mail) String() string {
	return "mail"
}
