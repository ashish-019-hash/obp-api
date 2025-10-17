package models

type BankAttribute struct {
	BankId          string            `json:"bank_id"`
	BankAttributeId string            `json:"bank_attribute_id"`
	Name            string            `json:"name"`
	AttributeType   BankAttributeType `json:"attribute_type"`
	Value           string            `json:"value"`
	IsActive        bool              `json:"is_active"`
}

func NewBankAttribute() *BankAttribute {
	return &BankAttribute{}
}
