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

func NewHandler(o Options) *mainHandler {
	return &mainHandler{
		Cfg:               o.Cfg,
		Log:               o.Log,
		ServiceToDoClient: o.ServiceToDoClient,
	}
}

type mainHandler struct {
	Cfg               config.Config
	Log               logger.Logger
	ServiceToDoClient service.ToDoClient
}

func (m *mainHandler) getAll(c *gin.Context) {
	var body genproto.Empty

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	res, err := m.ServiceToDoClient.ToDo().Get(ctx, &body)
	if err != nil {
		m.Log.Error("get method client", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (m *mainHandler) postEvent(c *gin.Context) {
	var body genproto.Event

	err := c.ShouldBindJSON(&body)

	if err != nil {
		m.Log.Error("Cant bind json", logger.Error(err))
		c.String(http.StatusBadRequest, "errror with json")
		return
	}

	ctx, finish := context.WithTimeout(context.Background(), time.Minute)

	defer finish()

	res, err := m.ServiceToDoClient.ToDo().Push(ctx, &body)

	if err != nil {
		m.Log.Error("push method client", logger.Error(err))
		c.String(http.StatusBadRequest, "error in Push function", err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (m *mainHandler) getByTime(c *gin.Context) {
	var timeReq genproto.Time

	timeReq.Time = c.Param("time")

	ctx, finish := context.WithTimeout(context.Background(), time.Minute)

	defer finish()

	res, err := m.ServiceToDoClient.ToDo().GetByTime(ctx, &timeReq)

	if err != nil {
		m.Log.Error("getbytime method client", logger.Error(err))
		c.String(http.StatusBadRequest, "not working getbytime function")
		return
	}

	c.JSON(http.StatusOK, res)
}

func (m *mainHandler) getById(c *gin.Context) {
	var id genproto.Id

	id.Id = c.Param("id")

	ctx, finish := context.WithTimeout(context.Background(), time.Minute)

	defer finish()

	res, err := m.ServiceToDoClient.ToDo().GetByID(ctx, &id)

	if err != nil {
		m.Log.Error("getbyid method client", logger.Error(err))
		c.String(http.StatusBadRequest, "getbyid doest work")
		return
	}

	c.JSON(http.StatusOK, res)

}

func (m *mainHandler) updateEvent(c *gin.Context) {
	var event genproto.Event

	err := c.ShouldBindJSON(&event)

	if err != nil {
		m.Log.Error("cant bind json updateEvent", logger.Error(err))
		c.String(http.StatusBadRequest, "invalid json to update")
		return
	}

	ctx, finish := context.WithTimeout(context.Background(), time.Minute)

	defer finish()

	res, err := m.ServiceToDoClient.ToDo().UpdateEvent(ctx, &event)

	if err != nil {
		m.Log.Error("cant get answer from service", logger.Error(err))
		c.String(http.StatusBadRequest, "error")
		return
	}

	c.JSON(http.StatusOK, res)
}

func (m *mainHandler) deleteEvent(c *gin.Context) {
	var id genproto.Id

	id.Id = c.Param("id")

	ctx, finish := context.WithTimeout(context.Background(), time.Minute)

	defer finish()

	res, err := m.ServiceToDoClient.ToDo().DeleteEvent(ctx, &id)

	if err != nil {
		m.Log.Error("delete", logger.Error(err))
		c.String(http.StatusBadRequest, "smt wrong")
		return
	}
	c.JSON(http.StatusOK, res)
}
