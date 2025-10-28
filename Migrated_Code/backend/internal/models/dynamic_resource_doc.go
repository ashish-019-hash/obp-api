package models

type DynamicResourceDoc struct {
	BankId                  string `json:"bank_id"`
	DynamicResourceDocId    string `json:"dynamic_resource_doc_id"`
	PartialFunctionName     string `json:"partial_function_name"`
	RequestVerb             string `json:"request_verb"`
	RequestUrl              string `json:"request_url"`
	Summary                 string `json:"summary"`
	Description             string `json:"description"`
	ExampleRequestBody      string `json:"example_request_body"`
	SuccessResponseBody     string `json:"success_response_body"`
	ErrorResponseBodies     string `json:"error_response_bodies"`
	Tags                    string `json:"tags"`
	Roles                   string `json:"roles"`
	MethodBody              string `json:"method_body"`
}

func NewDynamicResourceDoc() *DynamicResourceDoc {
	return &DynamicResourceDoc{}
}
