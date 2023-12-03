package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	user := "kiettran"
	pass := "quantrimang@2003"
	host := "testconnect.mysql.database.azure.com"
	db := "gotodo"

	credentials := fmt.Sprintf("%s:%s@(%s:3306)/%s?charset=utf8&parseTime=True", user, pass, host, db)

	database, err := sql.Open("mysql", credentials)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`USE gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	return database
}
