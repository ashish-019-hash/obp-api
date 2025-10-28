package models

import "time"

type Consent struct {
	ConsentId                        string    `json:"consent_id"`
	UserId                           string    `json:"user_id"`
	Secret                           string    `json:"secret"`
	Status                           string    `json:"status"`
	Challenge                        string    `json:"challenge"`
	JsonWebToken                     string    `json:"json_web_token"`
	ConsumerId                       string    `json:"consumer_id"`
	ConsentRequestId                 string    `json:"consent_request_id"`
	ApiStandard                      string    `json:"api_standard"`
	ApiVersion                       string    `json:"api_version"`
	RecurringIndicator               bool      `json:"recurring_indicator"`
	ValidUntil                       time.Time `json:"valid_until"`
	FrequencyPerDay                  int       `json:"frequency_per_day"`
	UsesSoFarTodayCounter            int       `json:"uses_so_far_today_counter"`
	UsesSoFarTodayCounterUpdatedAt   time.Time `json:"uses_so_far_today_counter_updated_at"`
	CombinedServiceIndicator         bool      `json:"combined_service_indicator"`
	LastActionDate                   time.Time `json:"last_action_date"`
	CreationDateTime                 time.Time `json:"creation_date_time"`
	StatusUpdateDateTime             time.Time `json:"status_update_date_time"`
	ExpirationDateTime               time.Time `json:"expiration_date_time"`
	TransactionFromDateTime          time.Time `json:"transaction_from_date_time"`
	TransactionToDateTime            time.Time `json:"transaction_to_date_time"`
	ConsentReferenceId               string    `json:"consent_reference_id"`
	Note                             string    `json:"note"`
}

func NewConsent() *Consent {
	return &Consent{}
}
