package handlers

import (
	"log/slog"
	"test-ozon/internal/http-server/middleware"
	"test-ozon/internal/service"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)


type Handlers struct{
	Service *service.Service
}

func NewHandlers(service *service.Service) *Handlers {
	return &Handlers{
		Service: service,
	}
}

func (h *Handlers) InitRouts(logger *slog.Logger) *gin.Engine{
	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(requestid.New())
	router.Use(middleware.LoggerMeddleware(logger))

	router.GET("/:alias", h.GetUrl)
	router.POST("/", h.SaveUrl)

	return router
}