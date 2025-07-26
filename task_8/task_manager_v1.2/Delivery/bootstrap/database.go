package bootstrap

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// refactored the database connector to make it testable
type MongoManager interface {
	Connect(ctx context.Context, uri string) (*mongo.Client, error)
	Ping(ctx context.Context, client *mongo.Client) error
	Disconnect(client *mongo.Client) error
}

type DefaultMongoManager struct{}

func (d *DefaultMongoManager) Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}

func (d *DefaultMongoManager) Ping(ctx context.Context, client *mongo.Client) error {
	return client.Ping(ctx, nil)
}

func (d *DefaultMongoManager) Disconnect(client *mongo.Client) error {
	return client.Disconnect(context.TODO())
}

func NewMongoDatabase(env *Env, mm MongoManager) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mm.Connect(ctx, env.MongoUri)
	if err != nil {
		return nil, err
	}

	if err := mm.Ping(ctx, client); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB.")
	return client, nil
}

func CloseMongoDBConnection(client *mongo.Client, mm MongoManager) error {
	if client == nil {
		return nil
	}

	err := mm.Disconnect(client)

	if err != nil {
		return  err
	}

	log.Println("Connection to MongoDB closed.")
	return  nil
}
