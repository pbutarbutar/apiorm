package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"pintekid/config"
	"pintekid/libs/security"
	"pintekid/schema/mutation"
	"pintekid/schema/query"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema, request *http.Request) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       context.WithValue(context.Background(), "token", request.URL.Query().Get("token")),
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	var Schame, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.RootQuery,
		Mutation: mutation.Mutation,
	})
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("Query"), Schame, r)
		json.NewEncoder(w).Encode(result)

	})

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphiql", graphiqlHandler)
	http.HandleFunc("/login", security.CreateTokenEndpoint)

	http.ListenAndServe(":"+config.PORT, nil)

}
