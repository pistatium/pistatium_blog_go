package main

import (
	"context"
	"os"
	"fmt"
	"time"
	"log"
	"net/http"
	"cloud.google.com/go/datastore"
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
	ctx := context.Background()
	projectID := os.Getenv("DATASTORE_PROJECT_ID")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q := datastore.NewQuery("Blog").Order("-datetime").Limit(10)

	// FIXME: ページング
	es := make([]Blog, 0, 10)
	if _, err := client.GetAll(ctx, q, &es); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v+", es)
}

func main() {
	http.HandleFunc("/", rootHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}
