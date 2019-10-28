package main

import (
	"os"
	"fmt"
	"time"
	"log"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Blog struct {
	Title      string
	Body       string
	More       string
	Category   string
	Datetime   time.Time
	Public     bool
	IsMarkdown bool
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Blog").Order("-datetime").Limit(10)
	es := make([]Blog, 0, 10)
	if _, err := q.GetAll(c, &es); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%v+", es)
}

func main() {
	http.HandleFunc("/", rootHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
