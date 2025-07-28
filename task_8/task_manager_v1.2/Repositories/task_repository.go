package repositories

import (
	"context"
	"fmt"
	domain "t8/taskmanager/Domain"
	"t8/taskmanager/Infrastructure/core/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

// wrapper
type taskRepository struct {
	database   mongo.Database
	collection string
}

// Add implements domain.TaskRepository.
func (tr *taskRepository) Add(ctx context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)
	_, error := collection.InsertOne(ctx, task)

	return error
}

// GetAll implements domain.TaskRepository.
func (t *taskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	collection := t.database.Collection(t.collection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetOne implements domain.TaskRepository.
func (t *taskRepository) GetOne(ctx context.Context, id string) (domain.Task, error) {
	collection := t.database.Collection(t.collection)

	var task domain.Task
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments() {
			return domain.Task{}, fmt.Errorf("task with ID %s not found", id)
		}
		return domain.Task{}, err
	}
	return task, nil
}

// Remove implements domain.TaskRepository.
func (t *taskRepository) Remove(ctx context.Context, id string) error {
	collection := t.database.Collection(t.collection)
	filter := bson.M{"_id": id}

	delete_count, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete task %w", err)
	}

	if delete_count == 0 {
		return fmt.Errorf("task with ID %s not found", id)
	}
	return nil
}

// Update implements domain.TaskRepository.
func (t *taskRepository) Update(ctx context.Context, id string, task *domain.Task) (domain.Task, error) {
	collection := t.database.Collection(t.collection)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"title":       task.Title,
		"description": task.Description,
		"due_date":    task.DueDate,
		"status":      task.Status,
	}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.Task{}, err
	}

	if result.MatchedCount == 0 {
		return domain.Task{}, fmt.Errorf("task with ID %s not found", id)
	}
	task.ID = id
	return *task, nil
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}
