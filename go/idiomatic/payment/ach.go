package payment

type AutomatedClearingHouse struct {
}

var _ Payment = &AutomatedClearingHouse{}

func (_ AutomatedClearingHouse) Kind() Kind {
	return Kinds.AutomatedClearningHouse
}
