package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	service    *Services
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func (e *Engine) NewError(err error) *Error {
	return &Error{
		service: e.Service,
		Message: err.Error(),
	}
}

func (e *Error) SetHttpCode(code int) *Error {
	e.StatusCode = code
	e.Status = http.StatusText(code)
	return e
}

func (e *Error) Send(c *gin.Context) {
	e.service.Logger.Info(e.service.Logger.WithFields(logrus.Fields{
		"error":       e.Message,
		"status_code": e.StatusCode,
	}))
	c.JSON(e.StatusCode, e)
	return
}
