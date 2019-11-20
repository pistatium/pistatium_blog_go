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

type Server struct {
	entries EntryRepo
}

type Photo struct {
	Id       int64      `datastore:"-"`
	Comment  string     `datastore:"comment"`
	Datetime *time.Time `datastore:"datetime"`
	Image    []byte     `datastore:"image"`
	Title    string     `datastore:"title"`
}

func (s *Server) getPhoto(gc *gin.Context) {
	ctx := gc.Request.Context()

	client, err = datastore.NewClient(ctx, )
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

func (s *Server) postEntry(gc *gin.Context) {

	ctx := gc.Request.Context()
	var entry Entry
	if err := gc.ShouldBindJSON(&entry); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.entries.CreateEntry(ctx, entry)

	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, entry)
}

func (s *Server) getEntry(gc *gin.Context) {

	ctx := gc.Request.Context()

	entryId := gc.Param("id")

	entry, err := s.entries.GetEntry(ctx, entryId)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ! entry.Public {
		gc.JSON(http.StatusForbidden, gin.H{"error": "private"})
	}
	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", CacheDuration))
	gc.JSON(http.StatusOK, &entry)
}

func (s *Server) getEntries(gc *gin.Context) {
	ctx := gc.Request.Context()

	page, _ := strconv.Atoi(gc.DefaultQuery("page", "0"))
	offset := 0
	if page != 0 {
		offset = page * EntriesPerPage
	}

	entries, err := s.entries.GetEntries(ctx, offset, EntriesPerPage, true)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", CacheDuration))
	gc.JSON(http.StatusOK, &Entries{Entries: entries})
}

func (s *Server) index(gc *gin.Context) {
	gc.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Pistatium Blog",
		"titleEnc":    "Pistatium Blog",
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
	projectID := os.Getenv(EnvKeyDatastoreProjectId) // Set by docker-compose
	if projectID == "" {
		projectID = os.Getenv(ProjectId) // Set by App Engine server
	}
	server := Server{
		entries: NewDatastoreEntryRepoImpl(projectID)
	}

	r := gin.Default()

	r.LoadHTMLGlob("front/dist/*.html")

	r.GET("/api/health", health)
	r.GET("/api/entries", server.getEntries)
	// FIXME LOGIN
	//r.POST("/api/entries", server.postEntry)
	r.GET("/api/entries/:id", server.getEntry)
	r.GET("/photo/show/:filename", server.getPhoto)
	r.NoRoute(server.index)
	log.Printf("Listening on port %s", port)
	entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(entryPoint)
}
