package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
)

func import_db() {
	db, err := sql.Open("mysql", "root:Admin@1234@tcp(127.0.0.1:3306)/performent_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO users(name, age, phone, email, address) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100000000; i++ {
		name := gofakeit.Name()
		age := gofakeit.Number(18, 65)
		phone := gofakeit.Phone()
		email := gofakeit.Email()
		address := gofakeit.Address().Address

		_, err = stmt.Exec(name, age, phone, email, address)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted row", i+1)
	}
}
