package currency

type Amount struct {
	Amount   float32  `json:"ammount"`
	Currency Currency `json:"currency"`
}
