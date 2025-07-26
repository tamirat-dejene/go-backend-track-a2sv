package bootstrap

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env   *Env
	Mongo *mongo.Client
}

func App(envPath string) Application {
	app := &Application{}
	env, err := NewEnv(envPath)
	if err != nil {
		log.Fatal(err)
	}

	mongo_client, err := NewMongoDatabase(app.Env, &DefaultMongoManager{})
	if err != nil {
		log.Fatal(err)
	}

	app.Mongo = mongo_client
	app.Env = env

	return *app
}

func (app *Application) CloseDBConnection() {
	err := CloseMongoDBConnection(app.Mongo, &DefaultMongoManager{})
	if err != nil {
		log.Println("Error closing MongoDB connection:", err)
	}
}
