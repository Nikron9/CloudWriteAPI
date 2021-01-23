package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	authModule "github.com/nikron9/CloudWriteAPI/auth"
)

func (db MongoDB) SignIn(username string, password string) (interface{}, error) {
	var err error
	var results UserEntity

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err = db.users.FindOne(ctx, q).Decode(&results); err != nil {
		return nil, fmt.Errorf("User not found.")
	}
	if err = authModule.CheckPassword(password, results.Password); err != nil {
		return nil, fmt.Errorf("Invalid password.")
	}
	results.Token, err = authModule.CreateToken(results.Username, results.Email)
	return results, nil
}

func (db MongoDB) SignUp(username string, password string, email string) (interface{}, error) {
	var err error
	var results UserEntity

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := db.users.FindOne(ctx, q).Decode(&results); err == nil {
		return nil, fmt.Errorf("User already exists.")
	}
	q = bson.M{"email": email}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	if err := db.users.FindOne(ctx, q).Decode(&results); err == nil {
		return nil, fmt.Errorf("Email already exists.")
	}
	pass, err := authModule.CipherPassword(password)
	if err != nil {
		return nil, err
	}
	results.ID = primitive.NewObjectID()
	results.Username = username
	results.Password = pass
	results.Email = email
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	_, err = db.users.InsertOne(ctx, results)
	if err != nil {
		return nil, err
	}
	results.Token, err = authModule.CreateToken(results.Username, results.Email)
	return results, nil
}
