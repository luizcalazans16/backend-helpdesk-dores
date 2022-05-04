package routes

import (
	"backend-helpdesk-dores/src/controller"
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
