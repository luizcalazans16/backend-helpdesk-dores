package routes

import (
	"backend-helpdesk-dores/src/controller"
	"net/http"
)

var routeLogin = Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	Function:              controller.Login,
	RequireAuthentication: false,
}
