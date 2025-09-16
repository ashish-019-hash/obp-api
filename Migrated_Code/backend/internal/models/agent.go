package models

type Agent struct {
	AgentId           string `json:"agent_id"`
	BankId            string `json:"bank_id"`
	LegalName         string `json:"legal_name"`
	MobileNumber      string `json:"mobile_number"`
	Email             string `json:"email"`
	IsConfirmedAgent  bool   `json:"is_confirmed_agent"`
}

func NewAgent() *Agent {
	return &Agent{}
}
