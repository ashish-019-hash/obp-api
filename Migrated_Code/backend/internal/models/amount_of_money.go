package models

type AmountOfMoney struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func NewAmountOfMoney() *AmountOfMoney {
	return &AmountOfMoney{}
}
