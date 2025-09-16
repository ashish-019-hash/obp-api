package models

type CustomerMessage struct {
	User           User   `json:"user"`
	BankId         string `json:"bank_id"`
	Message        string `json:"message"`
	FromDepartment string `json:"from_department"`
	FromPerson     string `json:"from_person"`
	Transport      string `json:"transport"`
}

func NewCustomerMessage() *CustomerMessage {
	return &CustomerMessage{}
}
