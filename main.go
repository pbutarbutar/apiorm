package main

import (
	"encoding/json"
	"fmt"
	"pintekid/schema/mutation"
	"pintekid/schema/query"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
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

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("Query"), Schame)
		json.NewEncoder(w).Encode(result)

	})
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":9000", nil)

}
