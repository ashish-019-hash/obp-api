package models

type CustomerAccountLink struct {
	CustomerAccountLinkId string `json:"customer_account_link_id"`
	CustomerId            string `json:"customer_id"`
	BankId                string `json:"bank_id"`
	AccountId             string `json:"account_id"`
	RelationshipType      string `json:"relationship_type"`
}

func NewCustomerAccountLink() *CustomerAccountLink {
	return &CustomerAccountLink{}
}
