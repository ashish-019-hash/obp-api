package repositories

import (
	"context"
	"database/sql"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type AgentRepository interface {
	Create(ctx context.Context, agent *models.Agent) error
	GetByID(ctx context.Context, id string) (*models.Agent, error)
	GetByBankID(ctx context.Context, bankID string) ([]*models.Agent, error)
	Update(ctx context.Context, id string, updateData map[string]interface{}) error
	Delete(ctx context.Context, id string) error
}

type agentRepository struct {
	db *sql.DB
}

func NewAgentRepository() AgentRepository {
	return &agentRepository{db: db.GetDB()}
}

func (r *agentRepository) Create(ctx context.Context, agent *models.Agent) error {
	return nil
}

func (r *agentRepository) GetByID(ctx context.Context, id string) (*models.Agent, error) {
	return &models.Agent{AgentId: id}, nil
}

func (r *agentRepository) GetByBankID(ctx context.Context, bankID string) ([]*models.Agent, error) {
	return []*models.Agent{}, nil
}

func (r *agentRepository) Update(ctx context.Context, id string, updateData map[string]interface{}) error {
	return nil
}

func (r *agentRepository) Delete(ctx context.Context, id string) error {
	return nil
}
