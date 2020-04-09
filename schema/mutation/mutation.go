package mutation

import (
	"pintekid/schema/types"

	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/*"CreateProducts": &graphql.Field{
			Type: graphql.NewList(types.ProductTypes),
			//config param argument
			Args: graphql.FieldConfigArgument{
				"ID_PRO": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"PRO_NAME": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"QTE_IN_STOCK": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"PRO_IMG": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: CreateProductMutation,
		},*/
		"SignupRegistrant": &graphql.Field{
			Type: types.UserRegistrantTypes,
			//config param argument
			Args: graphql.FieldConfigArgument{
				/*"registrant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},*/
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"fullname": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"phone_no": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: SignupRegistrantMutation,
		},
		// untuk membuat object lainya tinggal di ulang

	},
})
