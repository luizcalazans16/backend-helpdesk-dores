package main

import (
	"backend-helpdesk-dores/api"
	"backend-helpdesk-dores/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.CreateDatabase()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/artist", api.FindArtists)
	r.POST("/artist", api.CreateArtist)

	r.Run(":5000")
}
