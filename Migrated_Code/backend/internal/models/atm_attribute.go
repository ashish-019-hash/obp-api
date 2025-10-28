package models

type AtmAttribute struct {
	BankId          string           `json:"bank_id"`
	AtmId           string           `json:"atm_id"`
	AtmAttributeId  string           `json:"atm_attribute_id"`
	Name            string           `json:"name"`
	AttributeType   AtmAttributeType `json:"attribute_type"`
	Value           string           `json:"value"`
	IsActive        bool             `json:"is_active"`
}

func NewAtmAttribute() *AtmAttribute {
	return &AtmAttribute{}
}
