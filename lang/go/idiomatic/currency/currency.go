package currency

type Amount struct {
	Code   CurrencyCode `json:"code,omitempty"`
	Amount float32      `json:"amount,omitempty"`
}
