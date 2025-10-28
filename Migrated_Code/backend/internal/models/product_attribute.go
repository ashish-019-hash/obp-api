package models

type ProductAttribute struct {
	BankId              string               `json:"bank_id"`
	ProductCode         string               `json:"product_code"`
	ProductAttributeId  string               `json:"product_attribute_id"`
	Name                string               `json:"name"`
	AttributeType       ProductAttributeType `json:"attribute_type"`
	Value               string               `json:"value"`
	IsActive            bool                 `json:"is_active"`
}

func NewProductAttribute() *ProductAttribute {
	return &ProductAttribute{}
}
