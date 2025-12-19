package main

import (
	"github.com/SteliosKoulinas/recordStoreApi/db"
	"github.com/SteliosKoulinas/recordStoreApi/handlers"
	"github.com/SteliosKoulinas/recordStoreApi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.Album{}, &models.Users{})
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/albums", handlers.GetAlbums)
		api.POST("/albums", handlers.CreateAlbum)
		api.GET("/album/:id", handlers.GetAlbum)
		api.PUT("/album/:id", handlers.UpdateAlbum)
		api.DELETE("/album/:id", handlers.DeleteAlbum)
	}
	r.Run(":8080")
}
