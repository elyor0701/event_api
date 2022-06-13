package api

import (
	"evenApi/config"
	"evenApi/logger"
	"evenApi/service"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Cfg               config.Config
	Log               logger.Logger
	ServiceToDoClient service.ToDoClient
}

func New(o Options) *gin.Engine {
	route := gin.Default()

	v1 := NewHandler(o)

	route.GET("/get", v1.getAll)
	route.POST("/event", v1.postEvent)
	route.GET("/getbytime/:time", v1.getByTime)
	route.GET("/getbyid/:id", v1.getById)
	route.PUT("/updatevent", v1.updateEvent)
	route.DELETE("/delete/:id", v1.deleteEvent)

	return route
}
