package routes

import (
	"backend-helpdesk-dores/internal/core/controller"
	"net/http"
)

var routeLogin = Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	Function:              controller.Login,
	RequireAuthentication: false,
}
