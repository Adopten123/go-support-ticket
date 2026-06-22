package repository

import (
	"context"

	"github.com/Adopten123/go-support-ticket/internal/core/domain"
)

// ClientRepository описывает контракт для работы с базой данных клиентов.
type ClientRepository interface {
	Create(ctx context.Context, client *domain.Client) error
	GetByID(ctx context.Context, id int64) (*domain.Client, error)
	List(ctx context.Context, limit, offset int) ([]*domain.Client, error)
	Update(ctx context.Context, client *domain.Client) error
	Delete(ctx context.Context, id int64) error
}
