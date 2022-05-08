package routes

import (
	"backend-helpdesk-dores/internal/core/controller"
	"net/http"
)

var routeTicket = []Route{
	{
		URI:                   "/tickets",
		Method:                http.MethodPost,
		Function:              controller.NewTicket,
		RequireAuthentication: true,
	},
}
