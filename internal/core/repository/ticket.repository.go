package repository

import (
	"backend-helpdesk-dores/internal/core/model"
	"context"

	"github.com/google/uuid"
)

type TicketRepository interface {
	FindTicketById(ctx context.Context, id uuid.UUID) (*model.Ticket, error)
}
