package models

type User struct {
	UserName string `json:"user_name" binding:"required" bson:"user_name"`
	Password string `json:"password" binding:"required" bson:"password"`
}
