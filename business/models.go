package business

type ConvertMoney struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}