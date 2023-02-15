package transport

import (
	"example-goka/app/interface/container"
	"example-goka/app/usecase"
	"example-goka/app/usecase/request"
	"log"

	"github.com/gin-gonic/gin"
)

type transport struct {
	usecase usecase.Usecase
}

func NewTransport(container *container.Container) *transport {
	return &transport{
		usecase: container.Usecase,
	}
}

func (t *transport) Notif(c *gin.Context) {
	body := c.MustGet("body").(request.NotifRequest)
	check, err := t.usecase.Notif(&body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(check, "check")
	c.JSON(200, check)
}
