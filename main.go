package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-api-sample/database"
)

func rootHander(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("db connection failed")
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte("Hello World!!"))
}

func main() {
	http.HandleFunc("/", rootHander)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
