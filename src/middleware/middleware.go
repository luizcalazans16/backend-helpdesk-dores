package middleware

import (
	"backend-helpdesk-dores/src/authentication"
	"backend-helpdesk-dores/src/response"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidateToken(r); erro != nil {
			response.Error(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
