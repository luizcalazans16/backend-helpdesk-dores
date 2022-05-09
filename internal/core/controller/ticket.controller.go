package controller

import (
	"backend-helpdesk-dores/internal/core/service"
	"net/http"
)

type TicketController struct {
	ticketService *service.TicketService
}

func NewTicket(w http.ResponseWriter, r *http.Request) {

}
