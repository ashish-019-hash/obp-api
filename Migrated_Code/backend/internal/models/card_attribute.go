package models

type CardAttribute struct {
	CardId          string            `json:"card_id"`
	CardAttributeId string            `json:"card_attribute_id"`
	Name            string            `json:"name"`
	AttributeType   CardAttributeType `json:"attribute_type"`
	Value           string            `json:"value"`
}

func NewCardAttribute() *CardAttribute {
	return &CardAttribute{}
}
