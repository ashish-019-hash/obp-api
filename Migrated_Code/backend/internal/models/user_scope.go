package models

type UserScope struct {
	ScopeId string `json:"scope_id"`
	UserId  string `json:"user_id"`
}

func NewUserScope() *UserScope {
	return &UserScope{}
}
