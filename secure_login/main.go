package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/simpleLogin")
	if err != nil {
		log.Println("Error Connecting to Database")
	}

	http.HandleFunc("/", home)

	http.HandleFunc("/register", registerTempl)
	http.HandleFunc("/login", loginTempl)

	http.HandleFunc("/api/register", register)
	http.HandleFunc("/api/login", login)

	log.Println("Starting Listening on Port 443...")
	log.Println("Using self signed cert.")
	err = http.ListenAndServeTLS(":443", "ssl_config/server.crt", "ssl_config/server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
