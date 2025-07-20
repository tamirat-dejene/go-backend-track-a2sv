package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient    *mongo.Client
	TaskCollection *mongo.Collection
	UserCollection *mongo.Collection
)

func InitMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	MongoClient = client
	db := client.Database("taskdb")
	TaskCollection = db.Collection("tasks")
	UserCollection = db.Collection("users")

	// Seed only if collection is empty
	count, err := TaskCollection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return err
	}
	if count == 0 {
		if err := seedTasks(ctx); err != nil {
			return err
		}
	}

	return nil
}

func seedTasks(ctx context.Context) error {
	var taskDocs []interface{}
	for _, task := range Tasks {
		taskDocs = append(taskDocs, task)
	}
	_, err := TaskCollection.InsertMany(ctx, taskDocs)
	return err
}
