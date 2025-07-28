package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	domain "t8/taskmanager/Domain"
	mongo_mocks "t8/taskmanager/Infrastructure/core/database/mongo/mocks"
)

// Helper to create a repository with mocked DB and collection
func setup() (*mongo_mocks.MockDatabase, *mongo_mocks.MockCollection, domain.TaskRepository) {

	mockDB := new(mongo_mocks.MockDatabase)
	mockColl := new(mongo_mocks.MockCollection)
	mockDB.On("Collection", mock.Anything).Return(mockColl)

	repo := NewTaskRepository(mockDB, "tasks")
	return mockDB, mockColl, repo
}

// TestAddSuccess should add a task successfully
func TestAddSuccess(t *testing.T) {
	_, mockColl, repo := setup()

	task := &domain.Task{
		ID:          "1",
		Title:       "Test task",
		Description: "desc",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	mockColl.On("InsertOne", mock.Anything, task).Return(nil, nil).Once()

	err := repo.Add(context.Background(), task)

	assert.NoError(t, err)
	mockColl.AssertExpectations(t)
}

// TestAddFails should return an error when adding a task fails
func TestAddFails(t *testing.T) {
	_, mockCOll, repo := setup()

	task := &domain.Task{
		ID:          "1",
		Title:       "Test task",
		Description: "desc",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	mockErr := errors.New("insert failed")
	mockCOll.On("InsertOne", mock.Anything, task).Return(nil, mockErr).Once()

	err := repo.Add(context.Background(), task)

	assert.Error(t, err)
	assert.Equal(t, mockErr, err)
	mockCOll.AssertExpectations(t)
}

// TestGetAllSuccess should return all tasks successfully
func TestGetAllSuccess(t *testing.T) {
	_, mockColl, repo := setup()

	tasks := []domain.Task{
		{ID: "1", Title: "task1"},
		{ID: "2", Title: "task2"},
	}

	mockCursor := new(mongo_mocks.MockCursor)

	mockColl.On("Find", mock.Anything, bson.M{}).Return(mockCursor, nil).Once()
	mockCursor.On("Close", mock.Anything).Return(nil).Once()
	mockCursor.On("All", mock.Anything, mock.AnythingOfType("*[]domain.Task")).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*[]domain.Task)
		*arg = tasks
	}).Return(nil).Once()

	got, err := repo.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, tasks, got)
	mockColl.AssertExpectations(t)
	mockCursor.AssertExpectations(t)
}

// TestGetAllFails should return an error when fetching all tasks fails
func TestGetAllFails(t *testing.T) {
	_, mockColl, repo := setup()

	mockColl.On("Find", mock.Anything, bson.M{}).Return(nil, errors.New("find failed")).Once()

	_, err := repo.GetAll(context.Background())

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "find failed")
	mockColl.AssertExpectations(t)
}

// TestGetOneSuccess should return a task by ID successfully
func TestGetOneSuccess(t *testing.T) {
	_, mockColl, repo := setup()

	expectedTask := domain.Task{ID: "1", Title: "Test task"}

	mockSingleResult := new(mongo_mocks.MockSingleResult)
	mockColl.On("FindOne", mock.Anything, bson.M{"_id": "1"}).Return(mockSingleResult).Once()

	mockSingleResult.
		On("Decode", mock.AnythingOfType("*domain.Task")).
		Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.Task)
			*arg = expectedTask
		}).
		Return(nil).Once()

	got, err := repo.GetOne(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, got)
	mockColl.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

// TestGetOneNotFound should return an error when task is not found
func TestGetOneNotFound(t *testing.T) {
	_, mockColl, repo := setup()

	mockSingleResult := new(mongo_mocks.MockSingleResult)
	mockColl.On("FindOne", mock.Anything, bson.M{"_id": "1"}).Return(mockSingleResult).Once()
	mockSingleResult.On("Decode", mock.Anything).Return(mongo.ErrNoDocuments).Once()

	_, err := repo.GetOne(context.Background(), "1")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "task with ID 1 not found")
	mockColl.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

// TestGetOneError should return an error when decoding fails
func TestGetOneError(t *testing.T) {
	_, mockColl, repo := setup()

	mockSingleResult := new(mongo_mocks.MockSingleResult)
	mockColl.On("FindOne", mock.Anything, bson.M{"_id": "1"}).Return(mockSingleResult).Once()
	mockSingleResult.On("Decode", mock.Anything).Return(errors.New("decode error")).Once()

	_, err := repo.GetOne(context.Background(), "1")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decode error")
	mockColl.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

// TestRemoveSuccess should remove a task successfully
func TestRemoveSuccess(t *testing.T) {
	_, mockColl, repo := setup()

	mockColl.On("DeleteOne", mock.Anything, bson.M{"_id": "1"}).Return(int64(1), nil).Once()

	err := repo.Remove(context.Background(), "1")

	assert.NoError(t, err)
	mockColl.AssertExpectations(t)
}

// TestRemoveNotFound should return an error when task to remove is not found
func TestRemoveNotFound(t *testing.T) {
	_, mockColl, repo := setup()

	mockColl.On("DeleteOne", mock.Anything, bson.M{"_id": "1"}).Return(int64(0), nil).Once()

	err := repo.Remove(context.Background(), "1")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "task with ID 1 not found")
	mockColl.AssertExpectations(t)
}

// TestRemoveDeleteError should return an error when delete operation fails
func TestRemoveDeleteError(t *testing.T) {
	_, mockColl, repo := setup()

	mockColl.On("DeleteOne", mock.Anything, bson.M{"_id": "1"}).Return(int64(0), errors.New("db error")).Once()

	err := repo.Remove(context.Background(), "1")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to delete task")
	mockColl.AssertExpectations(t)
}

// TestUpdateSuccess should update a task successfully
func TestUpdateSuccess(t *testing.T) {
	_, mockColl, repo := setup()

	task := &domain.Task{
		Title:       "updated title",
		Description: "updated desc",
		DueDate:     time.Now(),
		Status:      "done",
	}

	mockColl.On("UpdateOne", mock.Anything, bson.M{"_id": "1"}, mock.Anything).Return(&mongo.UpdateResult{MatchedCount: 1}, nil).Once()

	updatedTask, err := repo.Update(context.Background(), "1", task)

	assert.NoError(t, err)
	assert.Equal(t, "1", updatedTask.ID)
	assert.Equal(t, task.Title, updatedTask.Title)
	mockColl.AssertExpectations(t)
}

// TestUpdateNotFound should return an error when task to update is not found
func TestUpdateNotFound(t *testing.T) {
	_, mockColl, repo := setup()

	task := &domain.Task{
		Title:       "updated title",
		Description: "updated desc",
		DueDate:     time.Now(),
		Status:      "done",
	}

	mockColl.On("UpdateOne", mock.Anything, bson.M{"_id": "1"}, mock.Anything).Return(&mongo.UpdateResult{MatchedCount: 0}, nil).Once()

	_, err := repo.Update(context.Background(), "1", task)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "task with ID 1 not found")
	mockColl.AssertExpectations(t)
}

// TestUpdateError should return an error when update operation fails
func TestUpdateError(t *testing.T) {
	_, mockColl, repo := setup()

	task := &domain.Task{
		Title:       "updated title",
		Description: "updated desc",
		DueDate:     time.Now(),
		Status:      "done",
	}

	mockColl.On("UpdateOne", mock.Anything, bson.M{"_id": "1"}, mock.Anything).Return(nil, errors.New("update error")).Once()

	_, err := repo.Update(context.Background(), "1", task)

	assert.Error(t, err)
	mockColl.AssertExpectations(t)
}
