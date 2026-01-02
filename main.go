package main

import (
	"github.com/SteliosKoulinas/recordStoreApi/db"
	"github.com/SteliosKoulinas/recordStoreApi/handlers"
	"github.com/SteliosKoulinas/recordStoreApi/middleware"
	"github.com/SteliosKoulinas/recordStoreApi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.Album{}, &models.Users{})
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		api.GET("/albums", handlers.GetAlbums)
		api.GET("/users", handlers.GetUsers)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			//protected.GET("/albums", handlers.GetAlbums)
			protected.POST("/albums", handlers.CreateAlbum)
			protected.GET("/album/:id", handlers.GetAlbum)
			protected.PUT("/album/:id", handlers.UpdateAlbum)
			protected.DELETE("/album/:id", handlers.DeleteAlbum)
		}
	}
	r.Run("127.0.0.1:3187")
}
