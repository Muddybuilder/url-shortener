package main

import (
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	switch r.Method {
	case http.MethodGet:
		queryURL := query.Get("short_url")
		// lookup the shortURL from db
		// if exist, return respond with original url with redit
		// else, return error

	case http.MethodPost:
		longURL := query.Get("long_url")
		// lookup the longURL from db
		// if exist, respond error
		// else, create permutation value

	}

	w.Header().Set("Content-Type", "application/json")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handleRequest)

	http.ListenAndServe(":8000", mux)

}
