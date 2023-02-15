package server

import (
	"example-goka/app/interface/container"
	"example-goka/app/interface/middleware"
	"example-goka/app/shared/config"
	"example-goka/app/transport"
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
