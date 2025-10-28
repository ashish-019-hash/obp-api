package models

type DynamicMessageDoc struct {
	BankId                   string `json:"bank_id"`
	DynamicMessageDocId      string `json:"dynamic_message_doc_id"`
	Process                  string `json:"process"`
	MessageFormat            string `json:"message_format"`
	Description              string `json:"description"`
	OutboundTopic            string `json:"outbound_topic"`
	InboundTopic             string `json:"inbound_topic"`
	ExampleOutboundMessage   string `json:"example_outbound_message"`
	ExampleInboundMessage    string `json:"example_inbound_message"`
	OutboundAvroSchema       string `json:"outbound_avro_schema"`
	InboundAvroSchema        string `json:"inbound_avro_schema"`
	AdapterImplementation    string `json:"adapter_implementation"`
	MethodBody               string `json:"method_body"`
	ProgrammingLang          string `json:"programming_lang"`
}

func NewDynamicMessageDoc() *DynamicMessageDoc {
	return &DynamicMessageDoc{}
}
