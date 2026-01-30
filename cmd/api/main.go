package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ron2112/gin_rest_api/internal/config"
	"github.com/ron2112/gin_rest_api/internal/database"
	"github.com/ron2112/gin_rest_api/internal/handlers"
	"github.com/ron2112/gin_rest_api/internal/middleware"
)

// migrate create -ext sql -dir migrations -seg create_todos_table

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration!", err)
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer pool.Close()

	router := gin.Default() // stores pointer of gin.Engine
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Todo api is running",
			"status":   "success",
			"database": "connected",
		})
	})

	// Auth routes
	router.POST("/auth/register", handlers.CreateUserHandler(pool))

	router.POST("/auth/login", handlers.LoginHandler(pool, cfg))

	// user routes
	router.GET("/users/email/:email", handlers.GetUserByEmailHandler(pool))

	router.GET("/users/id/:id", handlers.GetUserByIdHandler(pool))

	protected := router.Group("/todos")
	protected.Use(middleware.AuthMiddleWare(cfg))

	// todo routes

	{
		protected.POST("", handlers.CreateTodoHandler(pool))

		protected.GET("", handlers.GetAllTodosHandler(pool))

		protected.GET("/:id", handlers.GetTodoHAndler(pool))

		protected.PUT("/:id", handlers.UpdateTodoHandler(pool))

		protected.DELETE("/:id", handlers.DeleteTodoHandler(pool))
	}

	// Authorization Middleware test route
	router.GET("/auth/test", middleware.AuthMiddleWare(cfg), handlers.TestProtectedHandler())

	router.Run(":" + cfg.Port)
}
