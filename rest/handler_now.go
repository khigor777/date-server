package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khigor777/date-server/rest/response"
)

func (e *Engine) handlerNow(ctx *gin.Context) {
	time, err := e.Service.TimeService.GetFloat64()
	if err != nil {
		e.NewError(err).SetHttpCode(http.StatusBadRequest).Send(ctx)
		return
	}

	ctx.JSON(http.StatusOK, response.Now{
		Time: response.FloatNumber(time),
	})
}
