package models

type Bank struct {
	BankId             string `json:"bank_id"`
	ShortName          string `json:"short_name"`
	FullName           string `json:"full_name"`
	LogoUrl            string `json:"logo_url"`
	WebsiteUrl         string `json:"website_url"`
	SwiftBic           string `json:"swift_bic"`
	NationalIdentifier string `json:"national_identifier"`
	BankRoutingScheme  string `json:"bank_routing_scheme"`
	BankRoutingAddress string `json:"bank_routing_address"`
}

func NewBank() *Bank {
	return &Bank{}
}
