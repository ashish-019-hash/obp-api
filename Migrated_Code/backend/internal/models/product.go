package models

type Meta struct {
	License License `json:"license"`
}

type License struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	BankId                  string `json:"bank_id"`
	ProductCode             string `json:"product_code"`
	Name                    string `json:"name"`
	Category                string `json:"category"`
	Family                  string `json:"family"`
	SuperFamily             string `json:"super_family"`
	MoreInfoUrl             string `json:"more_info_url"`
	TermsAndConditionsUrl   string `json:"terms_and_conditions_url"`
	Description             string `json:"description"`
	Meta                    Meta   `json:"meta"`
}

func NewProduct() *Product {
	return &Product{}
}
