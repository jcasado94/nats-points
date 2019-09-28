package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcasado94/nats-points/handling"
)

func main() {
	h, err := handling.NewHandling()
	if err != nil {
		log.Panic("couldn't create handling")
	}
	r := mux.NewRouter()
	r.HandleFunc("/articlesTagged", h.HandleTagArticles).Methods("GET")
	r.HandleFunc("/articles", h.HandleArticles).Methods("GET")
	r.HandleFunc("/invalidate", h.HandleInvalidation).Methods("GET")
	r.HandleFunc("/info", h.HandleInformation).Methods("GET")

	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
