package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTitle(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllTitle(c *gin.Context) {

}

func (h *Handler) getTitleById(c *gin.Context) {

}

func (h *Handler) updateTitle(c *gin.Context) {

}

func (h *Handler) deleteTitle(c *gin.Context) {

}
