package routes

import (
	"github.com/ayserragm/backend-project/internal/config"
	"github.com/ayserragm/backend-project/internal/handlers"
	"github.com/ayserragm/backend-project/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, cfg *config.Config) {
	// Health
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK ✅"})
	})

	// Auth
	auth := handlers.NewAuthHandler(cfg)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/register", auth.Register)
		v1.POST("/auth/login", auth.Login)
	}

	// Users CRUD (mevcut handler’ların aynı kaldığını varsayıyorum)
	// NOT: Örnek olarak GET /users'i korumalı yapalım:
	protected := router.Group("/")
	protected.Use(middleware.AuthRequired(cfg))
	{
		protected.GET("/users", handlers.GetUsers)
		protected.POST("/users", handlers.CreateUser) // istersen açıkta da bırakabilirsin
		protected.DELETE("/users/:id", handlers.DeleteUser)

		// sadece admin’e özel örnek endpoint:
		admin := protected.Group("/admin")
		admin.Use(middleware.RequireRole("admin"))
		{
			admin.GET("/stats", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "admin stats ok"})
			})
		}
	}
}
