package models

type Entitlement struct {
	EntitlementId     string `json:"entitlement_id"`
	BankId            string `json:"bank_id"`
	UserId            string `json:"user_id"`
	RoleName          string `json:"role_name"`
	CreatedByProcess  string `json:"created_by_process"`
}

func NewEntitlement() *Entitlement {
	return &Entitlement{}
}
