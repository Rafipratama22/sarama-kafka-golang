package middleware

import (
	"example-goka/app/shared/config"
	"example-goka/app/usecase/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gitlab.mncbank.co.id/backend/log"
	"gitlab.mncbank.co.id/backend/response/v2"
)

type middleware struct {
	response response.ServerResponse
}

func NewMiddleware(response *response.ServerResponse) Middleware {
	return &middleware{
		response: *response,
	}
}

func (m *middleware) Notif(c *gin.Context) {
	var body request.NotifRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		log.Error(c, err)
		m.response.Error(c, http.StatusOK, config.RCMobeDataNotValid)
		c.Abort()
		return
	}

	v := validator.New()
	if err := v.Struct(body); err != nil {
		log.Error(c, err)
		m.response.Error(c, http.StatusOK, config.RCMobeDataNotValid)
		c.Abort()
		return
	}

	c.Set("body", body)
	c.Next()
}
