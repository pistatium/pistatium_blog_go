package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (s *Server) GetPhoto(gc *gin.Context) {
	ctx := gc.Request.Context()

	photoId := strings.Replace(gc.Param("filename"), ".jpg", "", 1)

	photo, err := s.Photos.GetPhoto(ctx, photoId)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", CacheDuration))
	gc.Data(http.StatusOK, "image/jpeg", photo.Image)
}
