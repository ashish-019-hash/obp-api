package models

import "time"

type EntitlementRequest struct {
	EntitlementRequestId string    `json:"entitlement_request_id"`
	BankId               string    `json:"bank_id"`
	User                 User      `json:"user"`
	RoleName             string    `json:"role_name"`
	Created              time.Time `json:"created"`
}

func NewEntitlementRequest() *EntitlementRequest {
	return &EntitlementRequest{}
}
