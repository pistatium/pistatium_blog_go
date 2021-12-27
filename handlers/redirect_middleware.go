package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const defaultAppEngineDomain = "kimihiro-n.appspot.com"
const redirectTo = "https://pistatium.dev"

func DomainRedirectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Host == defaultAppEngineDomain {
			c.Redirect(http.StatusPermanentRedirect, redirectTo + c.Request.URL.Path)
			return
		}
		c.Next()
	}
}
