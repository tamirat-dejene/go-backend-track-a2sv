package repositories

import (
	"context"
	"fmt"
	domain "t8/taskmanager/Domain"
	infrastructure "t8/taskmanager/Infrastructure"
	"t8/taskmanager/Infrastructure/constants"
	"t8/taskmanager/Infrastructure/core/database/mongo"


	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	database   mongo.Database
	collection string
	ps   *infrastructure.PasswordService
}

// Login implements domain.UserRepository.
func (u *userRepository) Login(ctx context.Context, user *domain.User) (*domain.User, error) {
	storedUser, err := u.FindOne(ctx, user.UserName)
	if err != nil && err.Error() == constants.NOT_FOUND {
		return nil, err
	}

	// check
	err = u.ps.ValidatePassword(storedUser.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("%s", constants.INVALID_CREDETNTIALS)
	}

	user.Password = "********"
	return user, nil
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) Delete(ctx context.Context, user_name string) error {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"user_name": user_name}

	del_cnt, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if del_cnt == 0 {
		return fmt.Errorf("%s", constants.NOT_FOUND)
	}
	return nil
}

// FindUser implements domain.UserRepository.
func (u *userRepository) FindOne(ctx context.Context, user_name string) (*domain.User, error) {
	var user domain.User
	collection := u.database.Collection(u.collection)

	filter := bson.M{"user_name": user_name}
	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, fmt.Errorf("%s", constants.NOT_FOUND)
	}
	return &user, nil
}

// GetAll implements domain.UserRepository.
func (u *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	collection := u.database.Collection(u.collection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []domain.User
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Register implements domain.UserRepository.
func (u *userRepository) Register(ctx context.Context, user *domain.User) (string, error) {
	collection := u.database.Collection(u.collection)
	_, err := u.FindOne(ctx, user.UserName)

	if err == nil {
		return "", fmt.Errorf("%s", constants.DUPLICATE_USERNAME)
	}

	hashedPassword, err := u.ps.GetHashedPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)
	_, err = collection.InsertOne(ctx, user)

	return user.UserName, err
}



func NewUserRepository(db mongo.Database, collection string, ps *infrastructure.PasswordService) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
		ps:         ps,
	}
}
