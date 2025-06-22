package main

import (
	"authApp/controllers"
	"authApp/middleware"
	"authApp/models"
	"log"
	"time"

	"github.com/gin-contrib/cors"
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

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000", // Добавьте явно ваш фронтенд-адрес
			"http://127.0.0.1:3000", // Без слеша в конце
			"http://localhost",
			"http://127.0.0.1",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

			// Защищенный роут для проверки токена
			protected.GET("/check", func(c *gin.Context) {
				c.JSON(200, gin.H{"status": "valid"})
			})

			// Защищенный роут для данных
			protected.GET("/data", func(c *gin.Context) {
				userID := c.MustGet("user_id").(uuid.UUID)
				c.JSON(200, gin.H{
					"message": "Secret data",
					"user_id": userID,
					"data":    []string{"item1", "item2", "item3"},
				})
			})
		}
	}

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
