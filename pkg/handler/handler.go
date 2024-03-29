package handler

import (
	"github.com/gin-gonic/gin"
	"titleNewsRestApi/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitHandler() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	title := router.Group("/title", h.userIdentity)
	{
		title.POST("/", h.createTitle)
		title.GET("/", h.getAllTitle)
		title.GET(":id", h.getTitleById)
		//title.PUT(":id", h.updateTitle)
		title.DELETE("/:id", h.deleteTitle)
	}

	return router
}
