package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khigor777/date-server/rest/response"
)

func (e *Engine) handlerNow(ctx *gin.Context) {
	time, _ := e.Service.TimeService.GetFloat64()
	ctx.JSON(http.StatusOK, response.Now{
		Time: time,
	})
}
