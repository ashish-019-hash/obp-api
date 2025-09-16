package models

type EndpointTag struct {
	EndpointTagId string `json:"endpoint_tag_id"`
	OperationId   string `json:"operation_id"`
	TagName       string `json:"tag_name"`
	BankId        string `json:"bank_id"`
}

func NewEndpointTag() *EndpointTag {
	return &EndpointTag{}
}
