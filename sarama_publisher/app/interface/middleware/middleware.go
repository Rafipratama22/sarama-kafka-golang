package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	Notif(c *gin.Context)
}
