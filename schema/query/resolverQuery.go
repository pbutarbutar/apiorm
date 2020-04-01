package query

import (
	"pintekid/config"
	"pintekid/schema/types"

	"github.com/graphql-go/graphql"
)

func ProductResolve(param graphql.ResolveParams) (interface{}, error) {
	var a types.Product
	var b []types.Product
	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}
	b = b[:0]

	result, err := db.Query("select ID_PRO,PRO_NAME,QTE_IN_STOCK,ifnull(PRO_IMG,'') from Products")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&a.IdPro, &a.ProName, &a.QteStock, &a.ProImg)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)

	}

	return b, nil
}

func ProductSingleResolve(param graphql.ResolveParams) (interface{}, error) {
	var a types.Product
	var b []types.Product
	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}
	b = b[:0]

	id_pro, _ := param.Args["ID_PRO"].(string)

	result, err := db.Query("select ID_PRO,PRO_NAME,QTE_IN_STOCK,ifnull(PRO_IMG,'') from Products where ID_PRO = ?", id_pro)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&a.IdPro, &a.ProName, &a.QteStock, &a.ProImg)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)

	}

	return b, nil
}
