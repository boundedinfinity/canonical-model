package payment

type Phone struct {
}

var _ Payment = &Phone{}

func (_ Phone) Kind() Kind {
	return Kinds.Phone
}
