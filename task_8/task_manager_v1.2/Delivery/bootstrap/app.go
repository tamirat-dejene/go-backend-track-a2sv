package bootstrap

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env   *Env
	Mongo *mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	mongo_client, err := NewMongoDatabase(app.Env, &DefaultMongoClientManager{})
	
	if err != nil {
		log.Fatal(err)
	}
	app.Mongo = mongo_client
	return *app
}

func (app *Application) CloseDBConnection() {
	err := CloseMongoDBConnection(app.Mongo, &DefaultMongoClientManager{})
	if err != nil {
		log.Println("Error closing MongoDB connection:", err)
	}
}
