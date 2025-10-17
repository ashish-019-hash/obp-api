package models

type AccountWebhook struct {
	AccountWebhookId string `json:"account_webhook_id"`
	BankId           string `json:"bank_id"`
	AccountId        string `json:"account_id"`
	TriggerName      string `json:"trigger_name"`
	Url              string `json:"url"`
	HttpMethod       string `json:"http_method"`
	HttpProtocol     string `json:"http_protocol"`
	CreatedByUserId  string `json:"created_by_user_id"`
	IsActive         bool   `json:"is_active"`
}

func NewAccountWebhook() *AccountWebhook {
	return &AccountWebhook{}
}
