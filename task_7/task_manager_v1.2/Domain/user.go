package domain

import (
	"context"
)

const (
	UserCollection = "users"
)

type User struct {
	// ID primitive.ObjectID `bson:"_id"`
	UserName string `json:"user_name" binding:"required" bson:"user_name"`
	Password string `json:"password" binding:"required" bson:"password"`
}

type UserRepository interface {
	FindOne(ctx context.Context, user_name string) (*User, error)
	Register(ctx context.Context, user *User) (string, error)
	Delete(ctx context.Context, user_name string) error
	GetAll(ctx context.Context) ([]User, error)
	Login(ctx context.Context, user *User) (*User, error)
}

type UserUsecase interface {
	FindOne(ctx context.Context, user_name string) (*User, error)
	Register(ctx context.Context, user *User) (string, error)
	Delete(ctx context.Context, user_name string) error
	GetAll(ctx context.Context) ([]User, error)
	Login(ctx context.Context, user *User) (*User, error)
}
