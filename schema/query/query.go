package query

import (
	"pintekid/schema/types"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"Products": &graphql.Field{
			Type: graphql.NewList(types.ProductTypes),
			//config param argument
			Args: graphql.FieldConfigArgument{
				"ID_PRO": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: ProductResolve,
		},
		"Productssingle": &graphql.Field{
			Type: graphql.NewList(types.ProductTypes),
			//config param argument
			Args: graphql.FieldConfigArgument{
				"ID_PRO": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: ProductSingleResolve,
		},
		// untuk membuat object lainya tinggal di ulang

	},
})
