package data

import (
	"context"
	"fmt"
	"t4/taskmanager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// In-memory task storage and service functions for managing tasks.

var Tasks = []models.Task{
	{
		ID:          "342c676e-2826-4055-adda-5d6867cd6a74",
		Title:       "Complete Go assignment",
		Description: "Finish the Go programming assignment for A2SV",
		DueDate:     time.Now().AddDate(0, 0, 2),
		Status:      models.ONGOING,
	},
	{
		ID:          "30525160-2907-4eea-afcc-cc454ee2808f",
		Title:       "Review pull requests",
		Description: "Review pending PRs on GitHub repository",
		DueDate:     time.Now().AddDate(0, 0, 1),
		Status:      models.DONE,
	},
	{
		ID:          "0009622d-ba8d-433d-8d13-023391d48872",
		Title:       "Prepare presentation",
		Description: "Prepare slides for the upcoming team meeting",
		DueDate:     time.Now().AddDate(0, 0, 5),
		Status:      models.ONGOING,
	},
	{
		ID:          "bc82d468-50fa-4e44-b997-1023a7a3d290",
		Title:       "Update documentation",
		Description: "Update the project documentation with recent changes",
		DueDate:     time.Now().AddDate(0, 0, 3),
		Status:      models.ONGOING,
	},
	{
		ID:          "92564d1c-f6a6-47f3-a9bb-3ab63e8b68e8",
		Title:       "Plan sprint retrospective",
		Description: "Organize and plan the upcoming sprint retrospective meeting",
		DueDate:     time.Now().AddDate(0, 0, 4),
		Status:      models.ONGOING,
	},
	{
		ID:          "a1b2c3d4-e5f6-7890-abcd-1234567890ef",
		Title:       "Conduct code review session",
		Description: "Schedule and conduct a code review session for the team",
		DueDate:     time.Now().AddDate(0, 0, 6),
		Status:      models.ONGOING,
	},
	{
		ID:          "f0e9d8c7-b6a5-4321-9876-abcdef123456",
		Title:       "Deploy new release",
		Description: "Deploy the latest release to the production environment",
		DueDate:     time.Now().AddDate(0, 0, 7),
		Status:      models.ONGOING,
	},
}

// Task service functions for managing tasks in memory.

func AddTask(task *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := TaskCollection.InsertOne(ctx, task)
	return err
}

func RemoveTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	res, err := TaskCollection.DeleteOne(ctx, filter)

	if err != nil {
		return fmt.Errorf("failed to delete task %w", err)
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("task with ID %s not found", id)
	}

	return nil
}

func UpdateTask(id string, u models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"title":       u.Title,
		"description": u.Description,
		"due_date":    u.DueDate,
		"status":      u.Status,
	}}

	result, err := TaskCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return models.Task{}, nil
	}

	if result.MatchedCount == 0 {
		return models.Task{}, fmt.Errorf("task with ID %s not found", id)
	}

	u.ID = id
	return u, nil
}

func GetTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTask(id string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task models.Task
	err := TaskCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, fmt.Errorf("task with ID %s not found", id)
		}
		return models.Task{}, err
	}
	return task, nil
}
