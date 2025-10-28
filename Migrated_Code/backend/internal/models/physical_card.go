package models

import "time"

type CardAction string

type PhysicalCard struct {
	CardId         string       `json:"card_id"`
	BankId         string       `json:"bank_id"`
	BankCardNumber string       `json:"bank_card_number"`
	CardType       string       `json:"card_type"`
	NameOnCard     string       `json:"name_on_card"`
	IssueNumber    string       `json:"issue_number"`
	SerialNumber   string       `json:"serial_number"`
	ValidFrom      time.Time    `json:"valid_from"`
	Expires        time.Time    `json:"expires"`
	Enabled        bool         `json:"enabled"`
	Cancelled      bool         `json:"cancelled"`
	OnHotList      bool         `json:"on_hot_list"`
	Technology     string       `json:"technology"`
	Networks       []string     `json:"networks"`
	Allows         []CardAction `json:"allows"`
	CustomerId     string       `json:"customer_id"`
	Cvv            *string      `json:"cvv,omitempty"`
	Brand          *string      `json:"brand,omitempty"`
}

func NewPhysicalCard() *PhysicalCard {
	return &PhysicalCard{
		Networks: make([]string, 0),
		Allows:   make([]CardAction, 0),
	}
}
