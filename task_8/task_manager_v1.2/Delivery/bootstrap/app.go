package bootstrap

import (
	"context"
	"log"
	"t8/taskmanager/Infrastructure/core/database/mongo"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App(envPath string) Application {
	app := &Application{}
	env, err := NewEnv(envPath)
	if err != nil {
		log.Fatal(err)
	}

	mongo_client, err := mongo.NewClient(env.MongoUri)
	if err != nil {
		log.Fatal(err)
	}

	app.Mongo = mongo_client
	app.Env = env

	return *app
}

func (app *Application) CloseDBConnection() {
	if app.Mongo != nil {
		err := app.Mongo.Disconnect(context.TODO())
		if err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}
}
