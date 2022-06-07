package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "user=postgres password=123456 host=localhost port=5432 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	//_, err = db.Exec("INSERT INTO \"users\"(username,password) VALUES ($1,$2);", "zzq12", "sit")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var username, password string
		var uid int
		err = rows.Scan(&uid, &username, &password)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(username, password)
	}
}
