package payment

type Venmo struct {
}

var _ Payment = &Venmo{}

func (_ Venmo) Kind() Kind {
	return Kinds.Venmo
}
