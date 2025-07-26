package bootstrap

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// refactored the database connector to make it testable
type MongoClientManager interface {
	Connect(ctx context.Context, uri string) (*mongo.Client, error)
	Ping(ctx context.Context, client *mongo.Client) error
	Disconnect(client *mongo.Client) error
}

type DefaultMongoClientManager struct{}

func (d *DefaultMongoClientManager) Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}

func (d *DefaultMongoClientManager) Ping(ctx context.Context, client *mongo.Client) error {
	return client.Ping(ctx, nil)
}

func (d *DefaultMongoClientManager) Disconnect(client *mongo.Client) error {
	return client.Disconnect(context.TODO())
}

func NewMongoDatabase(env *Env, mcm MongoClientManager) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mcm.Connect(ctx, env.MongoUri)
	if err != nil {
		return nil, err
	}

	if err := mcm.Ping(ctx, client); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB.")
	return client, nil
}

func CloseMongoDBConnection(client *mongo.Client, mcm MongoClientManager) error {
	if client == nil {
		return nil
	}

	err := mcm.Disconnect(client)

	if err != nil {
		return  err
	}

	log.Println("Connection to MongoDB closed.")
	return  nil
}
