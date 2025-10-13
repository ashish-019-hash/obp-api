package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ashish-019-hash/obp-api-backend/internal/services"
	"github.com/ashish-019-hash/obp-api-backend/internal/utils"
)

type MetricsController struct {
	orchestrationService *services.OrchestrationService
}

func NewMetricsController(orchestrationService *services.OrchestrationService) *MetricsController {
	return &MetricsController{
		orchestrationService: orchestrationService,
	}
}

type AggregateMetricsResponse struct {
	TotalCount       int     `json:"total_count"`
	AverageResponse  float64 `json:"average_response_time"`
	MinResponse      float64 `json:"min_response_time"`
	MaxResponse      float64 `json:"max_response_time"`
	TotalResponseTime float64 `json:"total_response_time"`
}

type MetricResponse struct {
	Date                     time.Time `json:"date"`
	URL                      string    `json:"url"`
	UserID                   string    `json:"user_id"`
	UserName                 string    `json:"user_name"`
	AppName                  string    `json:"app_name"`
	DeveloperEmail           string    `json:"developer_email"`
	ImplementedByPartialFunction string `json:"implemented_by_partial_function"`
	ImplementedInVersion     string    `json:"implemented_in_version"`
	ConsumerID               string    `json:"consumer_id"`
	Verb                     string    `json:"verb"`
	CorrelationID            string    `json:"correlation_id"`
	Duration                 int       `json:"duration"`
}

type MetricsResponse struct {
	Metrics []MetricResponse `json:"metrics"`
}

func (c *MetricsController) GetAggregateMetrics(ctx *gin.Context) {
	response := AggregateMetricsResponse{
		TotalCount:        100,
		AverageResponse:   250.5,
		MinResponse:       50.0,
		MaxResponse:       1000.0,
		TotalResponseTime: 25050.0,
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}

func (c *MetricsController) GetMetrics(ctx *gin.Context) {
	response := MetricsResponse{
		Metrics: []MetricResponse{
			{
				Date:                         time.Now(),
				URL:                          "/obp/v5.1.0/root",
				UserID:                       "user123",
				UserName:                     "testuser",
				AppName:                      "TestApp",
				DeveloperEmail:               "dev@example.com",
				ImplementedByPartialFunction: "root",
				ImplementedInVersion:         "v5.1.0",
				ConsumerID:                   "consumer123",
				Verb:                         "GET",
				CorrelationID:                "corr123",
				Duration:                     150,
			},
		},
	}

	utils.SendJSONResponse(ctx, http.StatusOK, response)
}
