package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pistatium/pistatium_blog_go/repos"
	"net/http"
	"strconv"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) GetAdminEntries(gc *gin.Context) {
	ctx := gc.Request.Context()

	page, _ := strconv.Atoi(gc.DefaultQuery("page", "0"))
	offset := 0
	if page != 0 {
		offset = page * EntriesPerPage
	}

	entries, err := s.Entries.GetEntries(ctx, offset, EntriesPerPage, false)
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

	err := s.Entries.CreateEntry(ctx, entry)

	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, entry)
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
