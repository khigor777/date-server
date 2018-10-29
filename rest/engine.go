package web

import "github.com/gin-gonic/gin"

type Engine struct {
	Router  *gin.Engine
	Service *Services
}

func NewEngine(service *Services) *Engine {
	return &Engine{
		Router:  gin.Default(),
		Service: service,
	}
}

func (e *Engine) Run(addr string) error {
	e.initDefaultMiddleware()
	e.initHandlers()
	return e.Router.Run(addr)
}

func (e *Engine) initHandlers() {

	timeGroup := e.Router.Group("/time")
	{
		timeGroup.GET("/now", e.handlerNow)
		//timeGroup.GET("/string", e.filterHandler)
		timeGroup.GET("/add", e.handlerAdd)
		//timeGroup.POST("/set", e.filterAddHandler)
		//timeGroup.POST("/reset", e.filterAddHandler)
	}
}

func (e *Engine) initDefaultMiddleware() {
	e.Router.Use(gin.Recovery())
}
