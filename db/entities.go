package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteEntity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `bson:"title" json:"name,omitempty"`
	Content     string             `bson:"content" json:"imageUrl,omitempty"`
	IsArchived  bool               `bson:"isArchived" json:"category,omitempty"`
	IsPrivate   bool               `bson:"isPrivate" json:"description,omitempty"`
	CreatedDate time.Time          `bson:"createdDate,omitempty" json:"createdDate,omitempty"`
	Username    string             `bson:"username,omitempty" json:"username,omitempty"`
}

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Username string             `bson:"username" json:"username,omitempty"`
	Password string             `bson:"password" json:"password,omitempty"`
	Email    string             `bson:"email" json:"email,omitempty"`
	Token    string             `json:"token,omitempty"`
}
