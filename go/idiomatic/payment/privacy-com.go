package payment

type PrivacyCom struct {
}

var _ Payment = &PrivacyCom{}

func (_ PrivacyCom) Kind() Kind {
	return Kinds.PrivacyCom
}
