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

type reqBody struct {
	Query string `json:"query"`
}

func processQuery(query string, schema graphql.Schema, request *http.Request) (result string) {
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       context.WithValue(context.Background(), "token", request.URL.Query().Get("token")),
	})

	if len(r.Errors) > 0 {
		fmt.Printf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)

	return fmt.Sprintf("%s", rJSON)

}

func gqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No query data", 400)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		var Schame, err2 = graphql.NewSchema(graphql.SchemaConfig{
			Query:    query.RootQuery,
			Mutation: mutation.Mutation,
		})
		if err2 != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "%s", processQuery(rBody.Query, Schame, r))

	})
}

func main() {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphql", gqlHandler())
	http.Handle("/graphiql", graphiqlHandler)
	http.HandleFunc("/login", security.CreateTokenEndpoint)
	fmt.Println("Server Runing port " + config.PORT)
	http.ListenAndServe(":"+config.PORT, nil)

}
