package data

import (
	"fmt"
	"t4/taskmanager/models"
	"time"

	"github.com/google/uuid"
)

var tasks = []models.Task{
	{
		ID:          uuid.New(),
		Title:       "Complete Go assignment",
		Description: "Finish the Go programming assignment for A2SV",
		DueDate:     time.Now().AddDate(0, 0, 2),
		Status:      models.ONGOING,
	},
	{
		ID:          uuid.New(),
		Title:       "Review pull requests",
		Description: "Review pending PRs on GitHub repository",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      models.DONE,
	},
	{
		ID:          uuid.New(),
		Title:       "Prepare presentation",
		Description: "Prepare slides for the upcoming team meeting",
		DueDate:     time.Now().AddDate(0, 0, 5),
		Status:      models.ONGOING,
	},
	{
		ID:          uuid.New(),
		Title:       "Update documentation",
		Description: "Update the project documentation with recent changes",
		DueDate:     time.Now().AddDate(0, 0, 3),
		Status:      models.ONGOING,
	},
	{
		ID:          uuid.New(),
		Title:       "Plan sprint retrospective",
		Description: "Organize and plan the upcoming sprint retrospective meeting",
		DueDate:     time.Now().AddDate(0, 0, 4),
		Status:      models.ONGOING,
	},
}

func AddTask(task *models.Task) error {
	tasks = append(tasks, *task)
	return nil
}

func RemoveTask(id uuid.UUID) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %s not found", id)
}

func UpdateTask(id uuid.UUID, u models.Task) (models.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			u.ID = id
			tasks[i] = u
			return u, nil
		}
	}
	return models.Task{}, fmt.Errorf("task with ID %s not found", id)
}

func GetTasks() ([]models.Task, error) {
	return tasks, nil
}

func GetTask(id uuid.UUID) (models.Task, error) {
	for _, t := range tasks {
		return t, nil
	}
	return models.Task{}, fmt.Errorf("task with ID %s not found", id)
}
