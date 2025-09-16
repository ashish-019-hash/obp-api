package models


type ATM struct {
	AtmId                   string    `json:"atm_id"`
	BankId                  string    `json:"bank_id"`
	Name                    string    `json:"name"`
	Address                 Address   `json:"address"`
	Location                Location  `json:"location"`
	IsAccessible            *bool     `json:"is_accessible,omitempty"`
	LocatedAt               *string   `json:"located_at,omitempty"`
	MoreInfo                *string   `json:"more_info,omitempty"`
	HasDepositCapability    *bool     `json:"has_deposit_capability,omitempty"`
	SupportedLanguages      *[]string `json:"supported_languages,omitempty"`
	Services                *[]string `json:"services,omitempty"`
	AccessibilityFeatures   *[]string `json:"accessibility_features,omitempty"`
	SupportedCurrencies     *[]string `json:"supported_currencies,omitempty"`
	MinimumWithdrawal       *string   `json:"minimum_withdrawal,omitempty"`
	Phone                   *string   `json:"phone,omitempty"`
}

func NewATM() *ATM {
	return &ATM{}
}
