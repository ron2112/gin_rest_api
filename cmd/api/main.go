package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ron2112/gin_rest_api/internal/config"
	"github.com/ron2112/gin_rest_api/internal/database"
	"github.com/ron2112/gin_rest_api/internal/handlers"
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
			"message": "Todo api is running",
			"status": "success",
			"database": "connected",
		})
	})

	router.POST("/todos", handlers.CreateTodoHandler(pool))

	router.GET("/todos", handlers.GetAllTodosHandler(pool))

	router.GET("/todos/:id", handlers.GetTodoHAndler(pool))

	router.Run(":" + cfg.Port)
}