package query

import (
	"pintekid/schema/types"

	"github.com/graphql-go/graphql"
)

const APPLICATION_INCOMPLETE_INCL = 1
const APPLICATION_COMPLETED_COMP = 2
const APPLICATION_TELEVERIFIED_VERY = 3
const APPLICATION_ANALYSING_ANAL = 4
const APPLICATION_REQUEST_ADDITIONAL_DATA_REQD = 5
const APPLICATION_ANALYSIS_COMPLETE_CRAF = 6
const APPLICATION_ANALYSIS_REJECTED_REJE = 7
const APPLICATION_APPROVED_APPR = 8
const APPLICATION_APPROVED_BY_LENDER_LAPP = 9
const APPLICATION_CONTRACT_CREATED_CMAK = 10
const APPLICATION_CONTRACT_SIGNED_CSIG = 11
const APPLICATION_DISBURSEMENT_PROPOSED_DREQ = 12
const APPLICATION_DISBURSED_DISB = 13
const APPLICATION_PAID_PAID = 14
const APPLICATION_CANCELLED_CANC = 15
const APPLICATION_DECLINED_AUTOMATICALLY_AURE = 16

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		/*"Products": &graphql.Field{
			Type: graphql.NewList(types.ProductTypes),
			//config param argument
			Args: graphql.FieldConfigArgument{
				"ID_PRO": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"limit": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
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
		},*/
		"EDUBorrowerLoan": &graphql.Field{
			Type: graphql.NewList(types.EduLoanAppTypes),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return EduLoanResolve(p, APPLICATION_APPROVED_APPR)
			},
		},
		"SMEBorrowerLoan": &graphql.Field{
			Type: graphql.NewList(types.SmeLoanAppTypes),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return SmeLoanResolve(p, APPLICATION_APPROVED_APPR)
			},
		},
		"EDUListHaveLender": &graphql.Field{
			Type: graphql.NewList(types.SmeLoanAppTypes),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return EduLoanResolve(p, APPLICATION_APPROVED_BY_LENDER_LAPP)
			},
		},
		"SMEListHaveLender": &graphql.Field{
			Type: graphql.NewList(types.SmeLoanAppTypes),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return SmeLoanResolve(p, APPLICATION_APPROVED_BY_LENDER_LAPP)
			},
		},
		// untuk membuat object lainya tinggal di ulang

	},
})
