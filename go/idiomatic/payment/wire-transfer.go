package payment

type WireTransfer struct {
}

var _ Payment = &WireTransfer{}

func (_ WireTransfer) Kind() Kind {
	return Kinds.WireTransfer
}
