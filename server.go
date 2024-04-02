package main

import (
	"fmt"
	"net/http"
)

var db = newDB()

func handleRequest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	switch r.Method {
	case http.MethodGet:
		queryURL := query.Get("short_url")
		if db.doesExist(queryURL) {
			fmt.Println("exist")
			msg := db.retrieve(queryURL)
			fmt.Fprintf(w, msg)

		} else {
			fmt.Println("Doesn't exist")
		}
		// lookup the shortURL from db
		// if exist, return respond with original url with redit
		// else, return error

	case http.MethodPost:
		longURL := query.Get("long_url")
		if db.doesExist(longURL) {
			fmt.Println("exist")
			msg := db.retrieve(longURL)
			fmt.Fprintf(w, msg)
		} else {
			fmt.Println("Doesn't exist")
			msg := db.insert(longURL)
			fmt.Fprintf(w, msg)

		}

		// lookup the longURL from db
		// if exist, respond error
		// else, create permutation value

	}

	w.Header().Set("Content-Type", "application/json")
}

func main() {
	fmt.Println("Start!")
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handleRequest)

	http.ListenAndServe(":8000", mux)

}
