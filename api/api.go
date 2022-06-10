package api

import (
	"context"
	"evenApi/config"
	"evenApi/genproto"
	"evenApi/logger"
	"evenApi/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Cfg               config.Config
	Log               logger.Logger
	ServiceToDoClient service.ToDoClient
}

func New(o Options) *gin.Engine {
	route := gin.Default()

	route.GET("/get", func(c *gin.Context) {
		var body genproto.Empty

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		res, err := o.ServiceToDoClient.ToDo().Get(ctx, &body)
		if err != nil {
			o.Log.Error("get method client", logger.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	route.POST("/event", func(c *gin.Context) {
		var body genproto.Event

		err := c.ShouldBindJSON(&body)

		if err != nil {
			o.Log.Error("Cant bind json", logger.Error(err))
			c.String(http.StatusBadRequest, "errror with json")
			return
		}

		ctx, finish := context.WithTimeout(context.Background(), time.Minute)

		defer finish()

		res, err := o.ServiceToDoClient.ToDo().Push(ctx, &body)

		if err != nil {
			o.Log.Error("push method client", logger.Error(err))
			c.String(http.StatusBadRequest, "error in Push function", err)
			return
		}

		c.JSON(http.StatusCreated, res)
	})

	route.GET("/getbytime/:time", func(c *gin.Context) {
		var timeReq genproto.Time

		timeReq.Time = c.Param("time")

		ctx, finish := context.WithTimeout(context.Background(), time.Minute)

		defer finish()

		res, err := o.ServiceToDoClient.ToDo().GetByTime(ctx, &timeReq)

		if err != nil {
			o.Log.Error("getbytime method client", logger.Error(err))
			c.String(http.StatusBadRequest, "not working getbytime function")
			return
		}

		c.JSON(http.StatusOK, res)
	})

	return route
}
