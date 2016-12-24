package blog

import (
	"fmt"
	"time"
	"net/http"

	"appengine"
	"appengine/datastore"
)

type Blog struct {
	Title string
	Body string
	More string
	Category string
	Datetime time.Time
	Public bool
	IsMarkdown bool
}

func init() {
	http.HandleFunc("/", rootHandler)
}

func entries(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Blog", "", 0, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Blog").Order("-datetime").Limit(10)
	es := make([]Blog, 0, 10)
	if _, err := q.GetAll(c, &es); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, "Hello, world!")
}