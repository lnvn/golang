package demo_http

import (
	"fmt"
	"log"
	"net/http"
)

func newPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is people handler")
	})
}

func newAnimalHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is animal handler")
	})
}

func MyMux() {
	mux := http.NewServeMux()
	mux.Handle("/people", newPeopleHandler())
	mux.Handle("/animal", newAnimalHandler())
	log.Fatal(http.ListenAndServe(":8080", mux))
}
