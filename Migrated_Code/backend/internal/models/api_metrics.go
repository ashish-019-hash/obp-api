package models

type APIMetrics struct {
	TotalCalls      int64   `json:"total_calls"`
	AverageDuration float64 `json:"average_duration"`
	MinDuration     int64   `json:"min_duration"`
	MaxDuration     int64   `json:"max_duration"`
	TotalDuration   int64   `json:"total_duration"`
}

type APIRanking struct {
	URL             string  `json:"url"`
	CallCount       int64   `json:"call_count"`
	AverageDuration float64 `json:"average_duration"`
	TotalDuration   int64   `json:"total_duration"`
	Rank            int     `json:"rank"`
}

type ConsumerRanking struct {
	ConsumerID          string  `json:"consumer_id"`
	TotalCalls          int64   `json:"total_calls"`
	UniqueAPIs          int     `json:"unique_apis"`
	AverageResponseTime float64 `json:"average_response_time"`
	TotalDuration       int64   `json:"total_duration"`
	Rank                int     `json:"rank"`
}

func NewAPIMetrics() *APIMetrics {
	return &APIMetrics{}
}

func NewAPIRanking() *APIRanking {
	return &APIRanking{}
}

func NewConsumerRanking() *ConsumerRanking {
	return &ConsumerRanking{}
}
