package models

type CustomerAttribute struct {
	CustomerId          string                `json:"customer_id"`
	CustomerAttributeId string                `json:"customer_attribute_id"`
	Name                string                `json:"name"`
	AttributeType       CustomerAttributeType `json:"attribute_type"`
	Value               string                `json:"value"`
}

func NewCustomerAttribute() *CustomerAttribute {
	return &CustomerAttribute{}
}
