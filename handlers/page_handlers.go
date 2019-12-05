package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Ellipsis(length int, text string) string {
	r := []rune(text)
	if len(r) > length {
		return string(r[0:length]) + "..."
	}
	return text
}

func (s *Server) Index(gc *gin.Context) {
	ctx := gc.Request.Context()
	path := gc.Request.URL.Path
	title := ""
	description := ""
	switch {
	case path == "/":
		title = "Top"
	case strings.HasPrefix(path, "/show/"):
		entryID := strings.Replace(path, "/show/", "", 1)
		entry, err := s.Entries.GetEntry(ctx, entryID)
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

func (s *Server) Sitemap(gc *gin.Context) {
	ctx := gc.Request.Context()
	sm := stm.NewSitemap(1)
	sm.SetDefaultHost("https://kimihiro-n.appspot.com")

	sm.Create()

	entries, err := s.Entries.GetEntries(ctx, 0, 1000, true)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sm.Add(stm.URL{{"loc", "/"}, {"changefreq", "daily"}})
	for _, entry := range entries {
		lastMod := entry.Datetime
		if entry.Updated != nil {
			lastMod = entry.Updated
		}
		sm.Add(stm.URL{
			{"loc", "/show/" + entry.Id},
			{"changefreq", "daily"},
			{"lastmod", lastMod.Format(time.RFC3339)},
		})
	}

	gc.Data(http.StatusOK, "text/xml", sm.XMLContent())
}
