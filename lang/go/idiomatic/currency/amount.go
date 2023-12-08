package currency

import "github.com/boundedinfinity/schema/idiomatic/id"

type CurrencyAmount struct {
	Type   CurrencyType `json:"type,omitempty"`
	Amount float32      `json:"amount,omitempty"`
}

var _ id.TypeNamer = &CurrencyAmount{}

func (t CurrencyAmount) TypeName() string {
	return id.TypeNamers.Dotted(CurrencyAmount{})
}

var CurrencyAmounts = currencyAmounts{}

type currencyAmounts struct{}

func (t currencyAmounts) Usd(amount float32) CurrencyAmount {
	return CurrencyAmount{
		Amount: amount,
		Type:   CurrencyTypes.USD,
	}
}

func (t currencyAmounts) Euro(amount float32) CurrencyAmount {
	return CurrencyAmount{
		Amount: amount,
		Type:   CurrencyTypes.EUR,
	}
}
