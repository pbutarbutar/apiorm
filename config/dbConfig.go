package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/*
	export DATABASE_HOST=localhost:3666
	export DATABASE_DBNAME=pintekid
	export DATABASE_USER=?
	export DATABASE_PASSWORD =?
*/

const PORT = "9001"

func GetConnection() (*sql.DB, error) {
	var DBHOST = os.Getenv("DATABASE_HOST")
	var DBNAME = os.Getenv("DATABASE_DBNAME")
	var DBUSER = os.Getenv("DATABASE_USER")
	var DBPASSWORD = os.Getenv("DATABASE_PASSWORD")

	db, err := sql.Open("mysql", DBUSER+":"+DBPASSWORD+"@tcp("+DBHOST+")/"+DBNAME)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
