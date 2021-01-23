package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Queries

func (db MongoDB) GetNotes(username string, searchTerm string, withArchived bool, onlyMine bool) (interface{}, error) {
	var results []NoteEntity
	var err error
	q := append([]bson.M{}, bson.M{})
	if onlyMine {
		q = append(q, bson.M{"username": username})
	} else {
		q = append(q, bson.M{"$or": []bson.M{{"isPrivate": false}, {"username": username}}})
	}
	if searchTerm != "" {
		q = append(q, bson.M{"$text": bson.M{"$search": searchTerm}})
	}
	if !withArchived {
		q = append(q, bson.M{"isArchived": false})
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.notes.Find(ctx, bson.M{"$and": q}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem NoteEntity
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

// getCurrentUser
func (db MongoDB) GetCurrentUser(username string) (interface{}, error) {
	var err error
	var results UserEntity

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.users.FindOne(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Mutations

func (db MongoDB) AddNote(username string, title string, content string, isPrivate bool) (interface{}, error) {
	var err error
	var results NoteEntity

	results.ID = primitive.NewObjectID()
	results.Title = title
	results.Content = content
	results.IsPrivate = isPrivate
	results.IsArchived = false
	results.CreatedDate = time.Now()
	results.Username = username
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_, err = db.notes.InsertOne(ctx, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db MongoDB) UpdateNote(_id string, user string, title string, content string, isPrivate bool, isArchived bool) (interface{}, error) {
	var err error
	var results NoteEntity

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id, "username": user}
	q2 := bson.M{"$set": bson.M{"title": title,
		"content":    content,
		"isPrivate":  isPrivate,
		"isArchived": isArchived,
	}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.notes.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}
	results.Title = title
	results.Content = content
	results.IsPrivate = isPrivate
	results.IsArchived = isArchived
	return results, nil
}
