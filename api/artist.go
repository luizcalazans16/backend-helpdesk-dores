package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Artist struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique;not null"`
}

type CreateArtistInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func FindArtists(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var artists []Artist
	db.Find(&artists)
	c.JSON(http.StatusOK, gin.H{"data": artists})
}

func CreateArtist(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CreateArtistInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	artist := Artist{Name: input.Name, Email: input.Email}
	db.Create(&artist)
	c.JSON(http.StatusOK, gin.H{"data": artist})
}
