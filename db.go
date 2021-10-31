package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect method is create mongo clint and context
func connect(uri string)(*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30 * time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// CreateNewMongoClient method is setting mongo db config information
func CreateNewMongoClient() (*mongo.Client, context.Context, error) {
	user := GetEnvVariableOrDefault("MONGO_USER", DEFAULT_MONGO_USER)
	password := GetEnvVariableOrDefault("MONGO_PASSWORD", DEFAULT_MONGO_PASSWORD)
	host := GetEnvVariableOrDefault("MONGO_HOST", DEFAULT_MONGO_HOST)
	database := GetEnvVariableOrDefault("MONGO_DATABASE", DEFAULT_MONGO_DATABASE)

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true", user, password, host, database)
	client, ctx, _ , err := connect(connStr)
	if err != nil {
		//log.Error(err)
		return client, nil, err
	}
	return client, ctx, nil
}
