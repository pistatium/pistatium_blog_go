package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const EntriesPerPage = 10

const (
	ProjectId                = "GOOGLE_CLOUD_PROJECT"
	EnvKeyDatastoreProjectId = "DATASTORE_PROJECT_ID"
	EnvKeyPORT               = "PORT"
	CacheDuration            = 60
)

type Entry struct {
	Id         int64      `datastore:"-"`
	Title      string     `datastore:"title,noindex"`
	Body       string     `datastore:"body,noindex"`
	More       string     `datastore:"more,noindex"`
	Category   string     `datastore:"category"`
	Datetime   *time.Time `datastore:"datetime"`
	Public     bool       `datastore:"public"`
	IsMarkdown bool       `datastore:"is_markdown,noindex"`
	//ModifyUser string            `datastore:"modify_user"`
	//CreateUser *datastore.Entity `datastore:"create_user"`
}

type Photo struct {
	Id int64 `datastore:"-"`
	Comment string `datastore:"comment"`
	Datetime *time.Time `datastore:"datetime"`
	Image []byte `datastore:"image"`
	Title string `datastore:"title"`
}

type Entries struct {
	Entries []*Entry `json:"entries"`
}

func getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	projectID := os.Getenv(EnvKeyDatastoreProjectId) // Set by docker-compose
	if projectID == "" {
		projectID = os.Getenv(ProjectId) // Set by App Engine server
	}
	client, err = datastore.NewClient(ctx, projectID)
	return
}

func getPhoto(gc *gin.Context) {
	ctx := context.Background()

	client, err := getDatastoreClient(ctx)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	filename := strings.Replace(gc.Param("filename"), ".jpg", "", 1)
	photoId, err := strconv.Atoi(filename)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	k := datastore.IDKey("Photo", int64(photoId), nil)
	e := new(Photo)
	err = client.Get(ctx, k, e)
	if err != nil && err != err.(*datastore.ErrFieldMismatch) {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", CacheDuration))
	gc.Data(http.StatusOK, "image/jpeg", e.Image)
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

func getEntry(gc *gin.Context) {
	ctx := context.Background()

	client, err := getDatastoreClient(ctx)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	entryId, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	k := datastore.IDKey("Blog", int64(entryId), nil)
	e := new(Entry)
	err = client.Get(ctx, k, e)
	if err != nil && err != err.(*datastore.ErrFieldMismatch) {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ! e.Public {
		gc.JSON(http.StatusForbidden, gin.H{"error": "private"})
	}
	e.Id = int64(entryId)
	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", CacheDuration))
	gc.JSON(http.StatusOK, &e)
}

func getEntries(gc *gin.Context) {
	ctx := context.Background()

	client, err := getDatastoreClient(ctx)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	page, _ := strconv.Atoi(gc.DefaultQuery("page", "0"))
	offset := 0
	if page != 0 {
		offset = page * EntriesPerPage
	}

	// 最新10件取得
	q := datastore.NewQuery("Blog").Filter("public =", true).Order("-datetime").Limit(EntriesPerPage).Offset(offset)
	entries := make([]*Entry, 0, 10)
	keys, err := client.GetAll(ctx, q, &entries)
	if err != nil && err != err.(*datastore.ErrFieldMismatch) {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, key := range keys {
		entries[i].Id = key.ID
	}
	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", CacheDuration))
	gc.JSON(http.StatusOK, &Entries{Entries: entries})
}

func index(gc *gin.Context) {
	gc.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Pistatium Blog",
		"titleEnc": "Pistatium Blog",
		"description": "",
	})
}

func health(gc *gin.Context) {
	gc.JSON(http.StatusOK, &map[string]string{"status": "ok",})
}

func main() {
	port := os.Getenv(EnvKeyPORT)
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.LoadHTMLGlob("front/dist/*.html")

	r.GET("/api/health", health)
	r.GET("/api/entries", getEntries)
	// FIXME LOGIN
	//r.POST("/api/entries", postEntry)
	r.GET("/api/entries/:id", getEntry)
	r.GET("/photo/show/:filename", getPhoto)
	r.NoRoute(index)

	log.Printf("Listening on port %s", port)
	entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(entryPoint)
}
