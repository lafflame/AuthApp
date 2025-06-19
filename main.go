package main

import (
	"authApp/controllers"
	"authApp/middleware"
	"authApp/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Подключение к PostgreSQL
	dsn := "host=localhost user=postgres dbname=authApp port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автомиграция
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Инициализация Gin
	r := gin.Default()

	// Инициализация контроллеров
	authController := controllers.NewAuthController(db)

	// Маршруты
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.POST("/refresh", authController.RefreshToken)
		}

		protected := api.Group("/protected")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/test", func(c *gin.Context) {
				userID := c.MustGet("user_id").(uuid.UUID)
				c.JSON(200, gin.H{"message": "Access granted", "user_id": userID})
			})
		}
	}

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
