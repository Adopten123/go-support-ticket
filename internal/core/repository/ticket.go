package repository

import (
	"context"

	"github.com/Adopten123/go-support-ticket/internal/core/domain"
)

type TicketRepository interface {
	Create(ctx context.Context, ticket *domain.Ticket) error
	GetByID(ctx context.Context, id int64) (*domain.Ticket, error)
	List(ctx context.Context, limit, offset int) ([]*domain.Ticket, error)
	ListByClientID(ctx context.Context, clientID int64, limit, offset int) ([]*domain.Ticket, error)
	Update(ctx context.Context, ticket *domain.Ticket) error
	Delete(ctx context.Context, id int64) error
}