package usecases

import (
	"context"
	"errors"
	domain "t8/taskmanager/Domain"
	"t8/taskmanager/Domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup() (*mocks.MockTaskRepository, *userUsecase) {
	mockRepo := new(mocks.MockUserRepository)
	mockUsecase := &userUsecase{
		userRepository: mockRepo,
		contextTimeout: 5 * time.Second,
	}

	return (*mocks.MockTaskRepository)(mockRepo), mockUsecase

}

// login should be success
func TestLoginSuccess(t *testing.T) {
	mockRepo, mockUseCase := setup()

	user := &domain.User{
		UserName: "testUserName",
		Password: "*****",
	}

	mockRepo.On("Login", mock.Anything, user).Return(user, nil)

	ruser, err := mockUseCase.Login(context.TODO(), user)

	assert.NoError(t, err)
	assert.Equal(t, ruser, user)
	mockRepo.AssertExpectations(t)

}

func TestLoginFails(t *testing.T) {
	mockRepo, mockUseCase := setup()

	user := &domain.User{
		UserName: "test",
		Password: "test",
	}

	mockRepo.On("Login", mock.Anything, user).Return(&domain.User{}, errors.New("Login fails"))

	usr, err := mockUseCase.Login(context.TODO(), user)
	
	assert.Error(t, err)
	assert.Empty(t, usr)
	assert.ErrorContains(t, err, "Login fails")
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	
	t.Run("err should be nil", func(t *testing.T) {
		mockRepo, mockUseCase := setup()
		mockRepo.On("Delete", mock.Anything, "test_username").Return(nil)
		err := mockUseCase.Delete(context.TODO(), "test_username")
		assert.Nil(t, err)
	})

	t.Run("err shouldn't be nil", func(t *testing.T) {
		mockRepo, mockUseCase := setup()
		mockRepo.On("Delete", mock.Anything, "test_username").Return(errors.New("user_name not found"))
		err := mockUseCase.Delete(context.TODO(), "test_username")
		assert.ErrorContains(t, err, "user_name not found")
	})
}

func TestFindOne(t *testing.T) {
	t.Run("Find one should be successful", func(t *testing.T) {
		mockRepo, mockUsecase := setup()

		usr := &domain.User {
			UserName: "test",
			Password: "test",
		}
		mockRepo.On("FindOne", mock.Anything, "user_name").Return(usr, nil)

		u, err := mockUsecase.FindOne(context.TODO(), "user_name")

		assert.Equal(t, usr, u)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindOne should fail and return empty user", func(t *testing.T) {
		mockRepo, mockUsecase := setup()
		
		mockRepo.On("FindOne", mock.Anything, "test").Return(&domain.User{}, errors.New("user not found"))

		u, err := mockUsecase.FindOne(context.TODO(), "test")

		assert.Empty(t, u)
		assert.ErrorContains(t, err, "user not found")
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll should return users successfully", func(t *testing.T) {
		mockRepo, mockUsecase := setup()
		users := []domain.User{
			{UserName: "user1", Password: "pass1"},
			{UserName: "user2", Password: "pass2"},
		}
		mockRepo.On("GetAll", mock.Anything).Return(users, nil)

		result, err := mockUsecase.GetAll(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, users, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAll should return error", func(t *testing.T) {
		mockRepo, mockUsecase := setup()
		mockRepo.On("GetAll", mock.Anything).Return([]domain.User{}, errors.New("db error"))

		result, err := mockUsecase.GetAll(context.TODO())

		assert.ErrorContains(t, err, "db error")
		assert.Empty(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	t.Run("Register should succeed", func(t *testing.T) {
		mockRepo, mockUsecase := setup()
		user := &domain.User{
			UserName: "newuser",
			Password: "newpass",
		}
		mockRepo.On("Register", mock.Anything, user).Return("new_id", nil)

		id, err := mockUsecase.Register(context.TODO(), user)

		assert.Equal(t, id, "new_id")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Register should fail", func(t *testing.T) {
		mockRepo, mockUsecase := setup()
		user := &domain.User{
			UserName: "failuser",
			Password: "failpass",
		}
		mockRepo.On("Register", mock.Anything, user).Return("", errors.New("registration error"))

		id, err := mockUsecase.Register(context.TODO(), user)

		assert.Empty(t, id)

		assert.ErrorContains(t, err, "registration error")
		mockRepo.AssertExpectations(t)
	})
}