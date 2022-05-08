package main

import (
	"backend-helpdesk-dores/internal/config"
	"backend-helpdesk-dores/internal/core/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	fmt.Println("API running")
	r := router.GenerateRoute()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
