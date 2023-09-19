package db

import (
	"context"

	"github.com/rajatxs/go-fconsole/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func MongoDb() *mongo.Database {
	return client.Database(config.MongoDbName())
}

func ConnectMongoDb(ctx context.Context) (err error) {
	clientOptions := options.Client().ApplyURI(config.MongoDbConnectionUrl())

	// connect to MongoDB
	if client, err = mongo.Connect(ctx, clientOptions); err != nil {
		return err
	}

	// check the connection
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	return nil
}

func DisconnectMongoDb(ctx context.Context) (err error) {
	if client == nil {
		return nil
	} else {
		return client.Disconnect(ctx)
	}
}
