package main

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/handlers"
	"todo-app/internal/models"
	"todo-app/pkg/database"
)

func main() {
	db := database.InitDB()

	database.Migrate(db, &models.Task{})

	router := gin.Default()

	handlers.RegisterRoutes(router, db)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
