package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"todo-app/internal/models"
	"todo-app/internal/services"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/tasks", getTasks(db))
	router.GET("/tasks/:id", getTaskById(db))
	router.POST("/tasks", createTask(db))
	router.PUT("/tasks/:id", updateTask(db))
	router.DELETE("/tasks/:id", deleteTask(db))
}

func getTasks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks := services.GetAllTasks(db)
		c.JSON(http.StatusOK, tasks)
	}
}

func createTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTask models.Task
		if err := c.Bind(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		services.AddTask(db, &newTask)
		c.JSON(http.StatusCreated, newTask)
	}
}

func getTaskById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		task, err := services.GetTaskById(db, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
		}
		c.JSON(http.StatusOK, task)
	}
}

func updateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updatedTask models.Task
		if err := c.BindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := services.UpdateTask(db, id, &updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, updatedTask)
	}
}

func deleteTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := services.DeleteTaskById(db, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	}
}
