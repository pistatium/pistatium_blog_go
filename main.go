package main

import (
	"fmt"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
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
	photos  PhotoRepo
	admin   AdminUserRepo
	conf    *Conf
}

func Ellipsis(length int, text string) string {
	r := []rune(text)
	if len(r) > length {
		return string(r[0:length]) + "..."
	}
	return text
}

func (s *Server) getPhoto(gc *gin.Context) {
	ctx := gc.Request.Context()

	photoId := strings.Replace(gc.Param("filename"), ".jpg", "", 1)

	photo, err := s.photos.GetPhoto(ctx, photoId)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", CacheDuration))
	gc.Data(http.StatusOK, "image/jpeg", photo.Image)
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
	ctx := gc.Request.Context()
	path := gc.Request.URL.Path
	title := ""
	description := ""
	switch {
	case path == "/":
		title = "Top"
	case strings.HasPrefix(path, "/show/"):
		entryID := strings.Replace(path, "/show/", "", 1)
		entry, err := s.entries.GetEntry(ctx, entryID)
		if err == nil {
			title = entry.Title
			description = Ellipsis(12, entry.Body)
		}
	default:
		title = "Pistatium Blog (" + path + ")"

	}
	params := map[string]string{
		"title":       title,
		"description": description,
		"titleEnc":    url.PathEscape(title),
	}
	gc.HTML(http.StatusOK, "index.html", params)
}

func (s *Server) sitemap(gc *gin.Context) {
	ctx := gc.Request.Context()
	sm := stm.NewSitemap(1)
	sm.SetDefaultHost("https://kimihiro-n.appspot.com")

	sm.Create()

	entries, err := s.entries.GetEntries(ctx, 0, 1000, true)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sm.Add(stm.URL{{"loc", "/"}, {"changefreq", "daily"}})
	for _, entry := range entries {
		sm.Add(stm.URL{{"loc", "/show/" + entry.Id}, {"changefreq", "daily"}})
	}

	gc.Data(http.StatusOK, "text/xml", sm.XMLContent())
}


type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) adminLogin(gc *gin.Context) {
	ctx := gc.Request.Context()
	var user LoginForm
	if err := gc.ShouldBindJSON(&user); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := s.admin.GetValidUser(ctx, user.Username, user.Password)
	if err != nil {
		if  err == err.(*LoginError) {
			gc.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		} else {
			gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	SetLogin(gc, u.Username)
	gc.JSON(http.StatusOK, &map[string]string{"status": "logged in",})
}


func health(gc *gin.Context) {
	gc.JSON(http.StatusOK, &map[string]string{"status": "ok",})
}


func main() {
	ctx := context.Background()
	port := os.Getenv(EnvKeyPORT)
	if port == "" {
		port = "8080"
	}
	projectID := os.Getenv(EnvKeyDatastoreProjectId) // Set by docker-compose
	if projectID == "" {
		projectID = os.Getenv(ProjectId) // Set by App Engine server
	}
	confRepo := NewDatastoreConfRepoImpl(projectID)

	conf, err := confRepo.GetConf(ctx)
	if err != nil {
		panic(err)
	}
	server := Server{
		entries: NewDatastoreEntryRepoImpl(projectID),
		photos:  NewDatastorePhotoRepoImpl(projectID),
		admin: NewAdminUserRepoImpl(),
		conf: conf,
	}

	r := gin.Default()

	store := cookie.NewStore([]byte(server.conf.Secret))
	r.Use(sessions.Sessions("SESSION", store))

	r.LoadHTMLGlob("front/dist/*.html")

	r.GET("/api/health", health)
	r.GET("/sitemap.xml", server.sitemap)
	r.GET("/api/entries", server.getEntries)
	// FIXME LOGIN
	//r.POST("/api/entries", server.postEntry)
	r.GET("/api/entries/:id", server.getEntry)
	r.GET("/photo/show/:filename", server.getPhoto)
	r.POST("/admin/login", server.adminLogin)
	r.NoRoute(server.index)
	log.Printf("Listening on port %s", port)
	entryPoint := fmt.Sprintf("0.0.0.0:%s", port)

	r.Run(entryPoint)
}
