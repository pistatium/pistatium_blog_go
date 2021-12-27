package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/pistatium/pistatium_blog_go/handlers"
	"github.com/pistatium/pistatium_blog_go/repos"
	"golang.org/x/net/context"
	"log"
	"os"
)

const (
	ProjectId                = "GOOGLE_CLOUD_PROJECT"
	EnvKeyDatastoreProjectId = "DATASTORE_PROJECT_ID"
	EnvKeyPORT               = "PORT"
	DefaultPort              = "8080"
)

func main() {
	ctx := context.Background()
	port := os.Getenv(EnvKeyPORT)
	if port == "" {
		port = DefaultPort
	}
	projectID := os.Getenv(EnvKeyDatastoreProjectId) // Set by docker-compose
	if projectID == "" {
		projectID = os.Getenv(ProjectId) // Set by App Engine server
	}
	confRepo := repos.NewDatastoreConfRepoImpl(projectID)

	conf, err := confRepo.GetConf(ctx)
	if err != nil {
		panic(err)
	}
	server := handlers.Server{
		Entries: repos.NewDatastoreEntryRepoImpl(projectID),
		Photos:  repos.NewDatastorePhotoRepoImpl(projectID),
		Admin:   repos.NewAdminUserRepoImpl(),
		Conf:    conf,
	}

	r := gin.Default()

	r.Use(handlers.DomainRedirectMiddleware())

	store := cookie.NewStore([]byte(server.Conf.Secret))
	r.Use(sessions.Sessions("SESSION", store))

	r.LoadHTMLGlob("front/dist/*.html")

	r.GET("/api/entries", server.GetEntries)

	r.GET("/api/entries/:id", server.GetEntry)
	r.GET("/photo/show/:filename", server.GetPhoto)

	adm := r.Group("/admin/api")
	adm.Use(handlers.LoginRequired())
	{
		adm.GET("/is_login", server.IsLogin)
		adm.GET("/entries", server.GetAdminEntries)
		adm.POST("/entries", server.PostEntry)
		adm.POST("/photos", server.UploadPhoto)
	}
	r.POST("/admin/api/login", server.AdminLogin)

	r.GET("/sitemap.xml", server.Sitemap)
	r.NoRoute(server.Index)

	log.Printf("Listening on port %s", port)
	entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	err = r.Run(entryPoint)
	panic(err)
}
