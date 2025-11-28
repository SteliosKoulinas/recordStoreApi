package handlers

import (
	"net/http"

	"github.com/SteliosKoulinas/recordStoreApi/db"
	"github.com/SteliosKoulinas/recordStoreApi/models"

	"github.com/gin-gonic/gin"
)

// GET all albums

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	db.DB.Find(&albums)
	c.JSON(200, albums)
}

// GET single album
func GetAlbum(c *gin.Context) {
	id := c.Param("id")

	var album models.Album
	result := db.DB.First(&album, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(200, album)
}

// create Album
func CreateAlbum(c *gin.Context) {
	var album models.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//if exists already
	var existing models.Album
	if err := db.DB.Where("artist = ? AND title = ?", album.Artist, album.Title).First(&existing).Error; err == nil {
		c.JSON(409, gin.H{"error": "Album already exists"})
		return
	}
	if err := db.DB.Create(&album).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, album)
}

// UPDATE album
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	var album models.Album
	if err := db.DB.First(&album, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}

	var input models.Album
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	album.Artist = input.Artist
	album.Title = input.Title
	album.Year = input.Year

	db.DB.Save(&album)

	c.JSON(200, album)
}

// DELETE album
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	var album models.Album
	if err := db.DB.First(&album, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}

	db.DB.Delete(&album)

	c.JSON(200, gin.H{"message": "Deleted"})
}
