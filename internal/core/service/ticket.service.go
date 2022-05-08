package service

import (
	"backend-helpdesk-dores/internal/core/model"
	"backend-helpdesk-dores/internal/core/repository"
	"context"

	"github.com/google/uuid"
)

type TicketService struct {
	ticketRepository repository.TicketRepository
}

func (t *TicketService) FindTicketById(ctx context.Context, id uuid.UUID) (*model.Ticket, error) {
	ticket, err := t.ticketRepository.FindTicketById(ctx, id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
