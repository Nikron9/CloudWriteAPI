package graph

import (
	"github.com/nikron9/CloudWriteAPI/graph/resolvers"
	"log"

	"github.com/graphql-go/graphql"
)

// Init the schema of GraphQL
func InitSchema() graphql.Schema {
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"notes": &graphql.Field{
					Type: graphql.NewList(NoteType),
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type:         graphql.String,
							DefaultValue: "",
						},
						"searchTerm": &graphql.ArgumentConfig{
							Type:         graphql.String,
							DefaultValue: "",
						},
						"withArchived": &graphql.ArgumentConfig{
							Type:         graphql.Boolean,
							DefaultValue: false,
						},
						"onlyMine": &graphql.ArgumentConfig{
							Type:         graphql.Boolean,
							DefaultValue: false,
						},
					},
					Resolve: resolvers.Notes,
				},
				"currentUser": &graphql.Field{
					Type:    UserType,
					Args:    graphql.FieldConfigArgument{},
					Resolve: resolvers.CurrentUser,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addNote": &graphql.Field{
					Type: NoteType,
					Args: graphql.FieldConfigArgument{
						"title": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"content": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"isPrivate": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Boolean),
						},
						"username": &graphql.ArgumentConfig{
							Type: graphql.String,
						}},
					Resolve: resolvers.AddNote,
				},
				"updateNote": &graphql.Field{
					Type: NoteType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"title": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"content": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"isPrivate": &graphql.ArgumentConfig{
							Type: graphql.Boolean,
						},
						"isArchived": &graphql.ArgumentConfig{
							Type: graphql.Boolean,
						}},
					Resolve: resolvers.UpdateNote,
				},
				"signIn": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: resolvers.SignIn,
				},
				"signUp": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: resolvers.SignUp,
				},
			},
		}),
		Types: []graphql.Type{ID},
	})
	if err != nil {
		log.Fatal(err)
	}
	return graphqlSchema

}
