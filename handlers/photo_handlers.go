package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const ImageCacheDuration = 60 *  60 * 24 * 30

func (s *Server) GetPhoto(gc *gin.Context) {
	ctx := gc.Request.Context()

	photoId := strings.Replace(gc.Param("filename"), ".jpg", "", 1)

	photo, err := s.Photos.GetPhoto(ctx, photoId)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	contentType := photo.ContentType
	if contentType == "" {
		contentType = "image/jpeg"
	}
	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", ImageCacheDuration))
	gc.Data(http.StatusOK, contentType, photo.Image)
}
