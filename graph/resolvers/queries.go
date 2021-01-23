package resolvers

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/nikron9/CloudWriteAPI/base"
	"github.com/nikron9/CloudWriteAPI/db"
)

func Notes(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	username := params.Context.Value("username").(string)
	if username == "" {
		return nil, base.NewHttpError(401, "Unauthorized")
	}

	searchTerm := params.Args["searchTerm"].(string)
	withArchived := params.Args["withArchived"].(bool)
	onlyMine := params.Args["onlyMine"].(bool)

	results, err = db.Db.GetNotes(username, searchTerm, withArchived, onlyMine)
	if err == nil {
		return nil, base.NewHttpError(500, fmt.Sprintf("Internal server error - %s", err))
	}

	return results, nil
}

func CurrentUser(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	username := params.Context.Value("username").(string)
	if username == "" {
		return nil, base.NewHttpError(401, "Unauthorized")
	}

	results, err = db.Db.GetCurrentUser(username)
	if err == nil {
		return nil, base.NewHttpError(500, fmt.Sprintf("Internal server error - %s", err))
	}

	return results, nil
}
