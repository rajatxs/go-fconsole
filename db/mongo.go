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

func ConnectMongoDb() (err error) {
	clientOptions := options.Client().ApplyURI(config.MongoDbConnectionUrl())

	// connect to MongoDB
	if client, err = mongo.Connect(context.Background(), clientOptions); err != nil {
		return err
	}

	// check the connection
	if err = client.Ping(context.Background(), nil); err != nil {
		return err
	}

	return nil
}

func DisconnectMongoDb() {

}
