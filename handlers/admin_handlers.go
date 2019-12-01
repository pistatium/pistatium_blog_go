package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pistatium/pistatium_blog_go/repos"
	"net/http"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
