package currency

type CurrencyType struct {
	Code CurrencyCode `json:"code,omitempty"`
	Name CurrencyName `json:"name,omitempty"`
}

var CurrencyTypes = currentTypes{
	USD: CurrencyType{
		Code: CurrencyCodes.USD,
		Name: CurrencyNames.UnitedStatesDollar,
	},
	EUR: CurrencyType{
		Code: CurrencyCodes.EUR,
		Name: CurrencyNames.Euro,
	},
}

type currentTypes struct {
	USD CurrencyType
	EUR CurrencyType
}
