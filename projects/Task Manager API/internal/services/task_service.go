package services

import (
	"task-manager/internal/models"
	"task-manager/internal/storage"
)

func CreateTask(title string) models.Task {
	task := models.Task{
		ID:        storage.NextID,
		Title:     title,
		Completed: false,
	}
	storage.Tasks = append(storage.Tasks, task)
	storage.NextID++
	return task
}

func GetTask() []models.Task {
	return storage.Tasks
}

func DeleteTask(id int) bool {
	for i, task := range storage.Tasks {
		if task.ID == id {
			storage.Tasks = append(storage.Tasks[:i], storage.Tasks[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateTask(id int, title string) (models.Task, bool) {
	for i, task := range storage.Tasks {
		if task.ID == id {
			storage.Tasks[i].Title = title
			return storage.Tasks[i], true
		}
	}
	return models.Task{}, false
}
