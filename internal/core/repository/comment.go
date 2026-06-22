package repository

import (
	"context"

	"github.com/Adopten123/go-support-ticket/internal/core/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) error
	ListByTicketID(ctx context.Context, ticketID int64) ([]*domain.Comment, error)
}
