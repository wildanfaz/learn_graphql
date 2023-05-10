package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	http.HandleFunc("/test", handler)

	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	root := graphql.ObjectConfig{Name: "root", Fields: fields}
	schemaCfg := graphql.SchemaConfig{Query: graphql.NewObject(root)}
	schema, err := graphql.NewSchema(schemaCfg)

	if err != nil {
		log.Fatal(err)
	}

	query := r.URL.Query().Get("query")

	params := graphql.Params{Schema: schema, RequestString: query}
	res := graphql.Do(params)

	json.NewEncoder(w).Encode(res)
}

var fields = graphql.Fields{
	"player": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "Player",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.ID,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						res := p.Source.(*playerData)
						return res.ID, nil
					},
				},
				"nickname": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						res := p.Source.(*playerData)
						return res.Nickname, nil
					},
				},
			},
		}),
		Resolve: Resolver,
	},
}

type playerData struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

func Resolver(p graphql.ResolveParams) (interface{}, error) {
	return &playerData{
		ID:       01,
		Nickname: "Flovvint",
	}, nil
}
