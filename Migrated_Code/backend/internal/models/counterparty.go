package models

type Counterparty struct {
	CounterpartyId             string `json:"counterparty_id"`
	CounterpartyName           string `json:"counterparty_name"`
	ThisBankId                 string `json:"this_bank_id"`
	ThisAccountId              string `json:"this_account_id"`
	OtherBankRoutingScheme     string `json:"other_bank_routing_scheme"`
	OtherBankRoutingAddress    string `json:"other_bank_routing_address"`
	OtherAccountRoutingScheme  string `json:"other_account_routing_scheme"`
	OtherAccountRoutingAddress string `json:"other_account_routing_address"`
	IsBeneficiary              bool   `json:"is_beneficiary"`
}

func NewCounterparty() *Counterparty {
	return &Counterparty{}
}
