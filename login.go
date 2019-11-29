package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const SessionUserIdKey = "user_id"

func loginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get(SessionUserIdKey)
		if userId == nil {
			c.AbortWithStatus(403)
		} else {
			c.Set(SessionUserIdKey, userId)
			c.Next()
		}
	}
}

func DoLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func SetLogin(c *gin.Context, UserId string) {
	session := sessions.Default(c)
	session.Set(SessionUserIdKey, UserId)
	session.Save()
}
