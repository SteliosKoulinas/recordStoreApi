package handlers

import (
	"github.com/SteliosKoulinas/recordStoreApi/db"
	"github.com/SteliosKoulinas/recordStoreApi/models"
	"github.com/SteliosKoulinas/recordStoreApi/services"

	"github.com/gin-gonic/gin"
)

// GET all albums

var albumService = services.AlbumService{}

func GetAlbums(c *gin.Context) {
	albums, err := albumService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, albums)
}

// GET single album
func GetAlbum(c *gin.Context) {
	id := c.Param("id")

	album, err := albumService.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(200, album)
}

// create Album

func CreateAlbum(c *gin.Context) {
	var input models.Album

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := albumService.Create(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, input)
}

// UPDATE album
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	// First: does this album exist?
	existing, err := albumService.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}

	// Bind JSON to a temp struct
	var input models.Album
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Update fields
	existing.Artist = input.Artist
	existing.Title = input.Title
	existing.Year = input.Year

	// Save changes
	if err := db.DB.Save(&existing).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, existing)
}

// DELETE album
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	err := albumService.Delete(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(404, gin.H{"error": "Album not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Album deleted"})
}
