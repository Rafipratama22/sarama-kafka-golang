package server

import (
	"example-sarama/app/interface/container"
	"example-sarama/app/interface/middleware"
	"example-sarama/app/shared/config"
	"example-sarama/app/transport"
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.mncbank.co.id/backend/log"
)

func StartServer(container container.Container) {
	app := gin.New()
	app.Use(gin.Recovery(), log.MiddlewareSecureV2)

	middle := middleware.NewMiddleware(&container.Response)
	trasnport := transport.SetupTransport(&container)
	setupRouter(trasnport, middle, app)

	fmt.Println(app.Run(config.Server.PORTHTTP))
}
