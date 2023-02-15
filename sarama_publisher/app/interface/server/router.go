package server

import (
	"example-sarama/app/interface/middleware"
	"example-sarama/app/transport"

	"github.com/gin-gonic/gin"
	"gitlab.mncbank.co.id/backend/lockey"
)

func setupRouter(t *transport.Tp, middle middleware.Middleware, app *gin.Engine) {
	v1 := app.Group("/v1/message").Use(lockey.DecryptBody)
	v1.POST("/notif", middle.Notif, t.Transport.Notif)
}
