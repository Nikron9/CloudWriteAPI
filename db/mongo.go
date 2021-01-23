package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	session *mongo.Client
	users   *mongo.Collection
	notes   *mongo.Collection
}

var Db MongoDB

func ModelConfig() {
	_, _ = Db.notes.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{"title": bsonx.String("text"), "content": bsonx.String("text")},
	})
}

func ConnectDb() MongoDB {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(DbConfig.mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return MongoDB{
		session: client,
		notes:   client.Database(DbConfig.mongoDb).Collection("notes"),
		users:   client.Database(DbConfig.mongoDb).Collection("users"),
	}
}

func (db MongoDB) CloseDb() {
	err := db.session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
