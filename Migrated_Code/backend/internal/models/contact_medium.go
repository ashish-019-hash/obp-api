package models

type ContactMedium struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewContactMedium() *ContactMedium {
	return &ContactMedium{}
}
