package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pistatium/pistatium_blog_go/repos"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const AdminEntriesPerPage = 30

func (s *Server) GetAdminEntries(gc *gin.Context) {
	ctx := gc.Request.Context()

	page, _ := strconv.Atoi(gc.DefaultQuery("page", "0"))
	offset := 0
	if page != 0 {
		offset = page * AdminEntriesPerPage
	}

	entries, err := s.Entries.GetEntries(ctx, offset, AdminEntriesPerPage, false)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, &repos.Entries{Entries: entries})
}

func (s *Server) PostEntry(gc *gin.Context) {

	ctx := gc.Request.Context()
	var entry repos.Entry
	if err := gc.ShouldBindJSON(&entry); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("id", entry.Id)
	err := s.Entries.UpdateEntry(ctx, entry.Id, entry)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, entry)
}

func (s *Server) UploadPhoto(gc *gin.Context) {
	ctx := gc.Request.Context()

	file, header , err := gc.Request.FormFile("file")
	filename := header.Filename
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	img := repos.Photo{
		Datetime: &now,
		Image:    bs,
		Title:    filename,
	}
	err = s.Photos.PutPhoto(ctx, &img)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, &map[string]string{"status": "uploaded", "id": strconv.Itoa(int(img.Id))})
}

func (s *Server) IsLogin(gc *gin.Context) {
	gc.JSON(http.StatusOK, &map[string]string{"status": "ok",})
}

func (s *Server) AdminLogin(gc *gin.Context) {
	ctx := gc.Request.Context()
	var user LoginForm
	if err := gc.ShouldBindJSON(&user); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := s.Admin.GetValidUser(ctx, user.Username, user.Password)
	if err != nil {
		if err == err.(*repos.LoginError) {
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
