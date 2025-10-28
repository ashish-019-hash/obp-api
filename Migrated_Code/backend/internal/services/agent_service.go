package services

import (
	"context"
	"obp-api-backend/internal/models"
	"obp-api-backend/internal/repositories"
)

type AgentService interface {
	CreateAgent(ctx context.Context, agent *models.Agent) error
	GetAgentsByBankID(ctx context.Context, bankID string) ([]*models.Agent, error)
	GetAgentByID(ctx context.Context, agentID string) (*models.Agent, error)
	UpdateAgentStatus(ctx context.Context, agentID string, updateData map[string]interface{}) error
}

type agentService struct {
	agentRepo repositories.AgentRepository
}

func NewAgentService(agentRepo repositories.AgentRepository) AgentService {
	return &agentService{
		agentRepo: agentRepo,
	}
}

func (s *agentService) CreateAgent(ctx context.Context, agent *models.Agent) error {
	return s.agentRepo.Create(ctx, agent)
}

func (s *agentService) GetAgentsByBankID(ctx context.Context, bankID string) ([]*models.Agent, error) {
	return s.agentRepo.GetByBankID(ctx, bankID)
}

func (s *agentService) GetAgentByID(ctx context.Context, agentID string) (*models.Agent, error) {
	return s.agentRepo.GetByID(ctx, agentID)
}

func (s *agentService) UpdateAgentStatus(ctx context.Context, agentID string, updateData map[string]interface{}) error {
	return s.agentRepo.Update(ctx, agentID, updateData)
}
