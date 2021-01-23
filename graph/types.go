package graph

import (
	"github.com/graphql-go/graphql"
)

var NoteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Note",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: ID,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"isArchived": &graphql.Field{
			Type: graphql.Boolean,
		},
		"isPrivate": &graphql.Field{
			Type: graphql.Boolean,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{ // graphql only
			Type: graphql.String,
		},
	},
})
