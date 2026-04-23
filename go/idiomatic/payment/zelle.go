package payment

type Zelle struct {
}

var _ Payment = &Zelle{}

func (_ Zelle) Kind() Kind {
	return Kinds.Zelle
}
