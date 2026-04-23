package payment

type WesternUnion struct {
}

var _ Payment = &WesternUnion{}

func (_ WesternUnion) Kind() Kind {
	return Kinds.WesternUnion
}
