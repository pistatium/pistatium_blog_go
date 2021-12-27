package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const defaultAppEngineDomain = "kimihiro-n.appspot.com"
const redirectTo = "https://pistatium.dev"

func DomainRedirectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("hostname req", c.Request.Host)
		println("hostname env", os.Getenv("HTTP_HOST"))
		if c.Request.Host == defaultAppEngineDomain {
			c.Redirect(http.StatusPermanentRedirect, redirectTo + c.Request.URL.Path)
			return
		}
		c.Next()
	}
}
