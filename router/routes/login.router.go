package routes

import (
	"backend-helpdesk-dores/controller"
	"net/http"
)

var routeLogin = Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	Function:              controller.Login,
	RequireAuthentication: false,
}
