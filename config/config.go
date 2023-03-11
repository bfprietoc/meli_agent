package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "admin"
	dbPass := "Meli2023"
	dbName := "meli"
	//dbHost := "(127.0.0.1:3306)"
	dbHostAws := "(melidbinstance-1.cnugeet3gyfb.us-east-2.rds.amazonaws.com)"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp"+dbHostAws+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
