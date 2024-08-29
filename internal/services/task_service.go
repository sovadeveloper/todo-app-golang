package services

import (
	"gorm.io/gorm"
	"todo-app/internal/models"
)

func GetAllTasks(db *gorm.DB) []models.Task {
	var tasks []models.Task
	db.Find(&tasks)
	return tasks
}

func AddTask(db *gorm.DB, task *models.Task) {
	db.Create(&task)
}

func GetTaskById(db *gorm.DB, id string) (models.Task, error) {
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		return task, err
	}
	return task, nil
}

func UpdateTask(db *gorm.DB, id string, updatedTask *models.Task) error {
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		return err
	}
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Completed = updatedTask.Completed
	db.Save(&task)
	return nil
}

func DeleteTaskById(db *gorm.DB, id string) error {
	if err := db.Delete(&models.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
