package main

import (
	"context"
	"os"
	"fmt"
	"time"
	"log"
	"net/http"
	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

const (
	ProjectId                = "GOOGLE_CLOUD_PROJECT"
	EnvKeyDatastoreProjectId = "DATASTORE_PROJECT_ID"
	EnvKeyPORT               = "PORT"
)

type Entry struct {
	Title      string     `datastore:"title,noindex`
	Body       string     `datastore:"body,noindex`
	More       string     `datastore:"more,noindex`
	Category   string     `datastore:"category`
	Datetime   *time.Time `datastore:"datetime`
	Public     bool       `datastore:"public`
	IsMarkdown bool       `datastore:"is_markdown,noindex`
}

type Entries struct {
	Entries []*Entry `json:"entries"`
}

func getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	projectID := os.Getenv(EnvKeyDatastoreProjectId)  // Set by docker-compose
	if projectID == "" {
		projectID = os.Getenv(ProjectId)  // Set by App Engine server
	}
	client, err = datastore.NewClient(ctx, projectID)
	return
}

func postEntry(gc *gin.Context) {
	ctx := context.Background()

	var entry Entry
	if err := gc.ShouldBindJSON(&entry); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := getDatastoreClient(ctx)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if entry.Datetime == nil {
		now := time.Now()
		entry.Datetime = &now
	}
	key := datastore.IncompleteKey("Blog", nil)
	if _, err := client.Put(ctx, key, &entry); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, entry)
}

func getEntries(gc *gin.Context) {
	ctx := context.Background()

	client, err := getDatastoreClient(ctx)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 最新10件取得
	q := datastore.NewQuery("Blog").Order("-datetime").Limit(10)
	entries := make([]*Entry, 0, 10)
	if _, err := client.GetAll(ctx, q, &entries); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, &Entries{Entries: entries})
}

func index(gc *gin.Context) {
	gc.String(http.StatusOK, "try: GET /entries or POST /entries")
}

func main() {
	port := os.Getenv(EnvKeyPORT)
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/", index)
	r.GET("/entries", getEntries)
	r.POST("/entries", postEntry)

	log.Printf("Listening on port %s", port)
	entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(entryPoint)
}
