package resolvers

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/nikron9/CloudWriteAPI/base"
	"github.com/nikron9/CloudWriteAPI/db"
)

func AddNote(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	username := params.Context.Value("username").(string)
	if username == "" {
		return nil, base.NewHttpError(401, "Unauthorized")
	}

	title := params.Args["title"].(string)
	content := params.Args["content"].(string)
	isPrivate := params.Args["isPrivate"].(bool)

	results, err = db.Db.AddNote(username, title, content, isPrivate)
	if err != nil {
		return nil, base.NewHttpError(500, fmt.Sprintf("Internal server error - %s", err))
	}

	return results, nil
}

func UpdateNote(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	username := params.Context.Value("username").(string)
	if username == "" {
		return nil, base.NewHttpError(401, "Unauthorized")
	}

	id := params.Args["_id"].(string)
	title := params.Args["title"].(string)
	content := params.Args["content"].(string)
	isPrivate := params.Args["isPrivate"].(bool)
	isArchived := params.Args["isArchived"].(bool)

	results, err = db.Db.UpdateNote(id, username, title, content, isPrivate, isArchived)
	if err == nil {
		return nil, base.NewHttpError(500, fmt.Sprintf("Internal server error - %s", err))
	}

	return results, nil
}

func SignIn(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Args["username"].(string)
	pass := p.Args["password"].(string)

	results, err = db.Db.SignIn(user, pass)
	if err == nil {
		return nil, base.NewHttpError(500, fmt.Sprintf("Internal server error - %s", err))
	}

	return results, nil
}

func SignUp(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	user := p.Args["username"].(string)
	pass := p.Args["password"].(string)
	email := p.Args["email"].(string)

	results, err = db.Db.SignUp(user, pass, email)
	if err == nil {
		return nil, base.NewHttpError(500, fmt.Sprintf("Internal server error - %s", err))
	}

	return results, nil
}