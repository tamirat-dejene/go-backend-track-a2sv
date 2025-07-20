package data

import (
	"context"
	"fmt"
	"os"
	"t4/taskmanager/constants"
	"t4/taskmanager/models"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var (
	ATS = os.Getenv("ACCESS_TOKEN_SECRET")
	RTS = os.Getenv("REFRESH_TOKEN_SECRET")
	ATE = os.Getenv("ACCESS_TOKEN_EXPIRY")
	RTE = os.Getenv("REFRESH_TOKEN_EXPIRY")
)

type JWTPayload struct {
	UserName string
	Exp      string
}

func SignUser(payload *JWTPayload, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_name": payload.UserName,
		"exp":       payload.Exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(token_str string, secret_key []byte) (string, error) {
	token, err := jwt.Parse(token_str, func(t *jwt.Token) (any, error) {
		return secret_key, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	user_name := claims["user_name"].(string)
	return user_name, nil
}

func findUser(user_name string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{"user_name": user_name}
	err := UserCollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, fmt.Errorf("%s", constants.NOT_FOUND)
	}
	return &user, nil
}

func LoginUser(user *models.User) (*models.User, error) {
	storedUser, err := findUser(user.UserName)
	if err != nil && err.Error() == constants.NOT_FOUND {
		return nil, err
	}

	// check
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("%s", constants.INVALID_CREDETNTIALS)
	}

	user.Password = "********"
	return user, nil
}

func RegisterUser(user *models.User) (string, error) {
	_, err := findUser(user.UserName)

	if err == nil {
		return "", fmt.Errorf("%s", constants.DUPLICATE_USERNAME)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// salt and hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	_, err = UserCollection.InsertOne(ctx, user)

	return user.UserName, err
}

func DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_name": id}
	res, err := UserCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("%s", constants.NOT_FOUND)
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
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
