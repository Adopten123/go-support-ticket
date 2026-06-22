package repository

import (
	"context"

	"github.com/Adopten123/go-support-ticket/internal/core/domain"
)

type AgentRepository interface {
	Create(ctx context.Context, agent *domain.Agent) error
	GetByID(ctx context.Context, id int64) (*domain.Agent, error)
	List(ctx context.Context, limit, offset int) ([]*domain.Agent, error)
}