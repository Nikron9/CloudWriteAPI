package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/nikron9/CloudWriteAPI/auth"
	"github.com/nikron9/CloudWriteAPI/db"
	"github.com/nikron9/CloudWriteAPI/graph"
)

func main() {
	db.Db = db.ConnectDb()
	db.ModelConfig()
	defer db.Db.CloseDb()

	schema := graph.InitSchema()
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", auth.DisableCors(auth.AddAuthHeader(h)))
	err := http.ListenAndServe(serverConfig.serverUri, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
