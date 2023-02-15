package transport

import (
	"example-goka/app/usecase"

	"github.com/gin-gonic/gin"
)

type transport struct {
	usecase usecase.Usecase
}

func NewTransport(usecase usecase.Usecase) *transport {
	return &transport{
		usecase: usecase,
	}
}

func (t *transport) Notif(c *gin.Context) {
	c.JSON(200, gin.H{
		"check": "yess",
	})
}
