package main

import (
	"fmt"
	"os"
	"time"

	"github.com/SteliosKoulinas/recordStoreApi/db"
	"github.com/SteliosKoulinas/recordStoreApi/handlers"
	"github.com/SteliosKoulinas/recordStoreApi/middleware"
	"github.com/SteliosKoulinas/recordStoreApi/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

func main() {
	// Load .env file (optional, won't fail if not found)
	_ = godotenv.Load()

	db.Connect()
	db.DB.AutoMigrate(&models.Album{}, &models.Users{})
	r := gin.Default()

	authLimiter := middleware.NewIPRateLimiter(rate.Every(time.Minute/5), 5)     // 5 req/min
	publicLimiter := middleware.NewIPRateLimiter(rate.Every(time.Minute/60), 20) // 60 req/min

	api := r.Group("/api")
	{
		auth := api.Group("/")
		auth.Use(middleware.RateLimitMiddleware(authLimiter))
		{
			api.POST("/register", handlers.Register)
			api.POST("/login", handlers.Login)
		}
		public := api.Group("/")
		public.Use(middleware.RateLimitMiddleware(publicLimiter))
		{
			api.GET("/albums", handlers.GetAlbums)
			api.GET("/users", handlers.GetUsers)
		}
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		protected.Use(middleware.RateLimitMiddleware(publicLimiter))
		{
			//protected.GET("/albums", handlers.GetAlbums)
			protected.POST("/albums", handlers.CreateAlbum)
			protected.GET("/album/:id", handlers.GetAlbum)
			protected.PUT("/album/:id", handlers.UpdateAlbum)
			protected.DELETE("/album/:id", handlers.DeleteAlbum)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3187"
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("Starting server on %s\n", addr)
	r.Run(addr)
}
