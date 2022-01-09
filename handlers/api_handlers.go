package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pistatium/pistatium_blog_go/repos"
	"net/http"
	"strconv"
)

const (
	CacheDuration  = 60 * 10
	EntriesPerPage = 10
)

func (s *Server) GetEntry(gc *gin.Context) {

	ctx := gc.Request.Context()

	entryId := gc.Param("id")

	entry, err := s.Entries.GetEntry(ctx, entryId)
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

func (s *Server) GetEntries(gc *gin.Context) {
	ctx := gc.Request.Context()
	page, _ := strconv.Atoi(gc.DefaultQuery("page", "0"))
	offset := 0
	if page != 0 {
		offset = page * EntriesPerPage
	}

	entries, err := s.Entries.GetEntries(ctx, offset, EntriesPerPage, true)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", CacheDuration))
	gc.JSON(http.StatusOK, &repos.Entries{Entries: entries})
}
