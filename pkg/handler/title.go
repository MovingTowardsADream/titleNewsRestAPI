package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	news "titleNewsRestApi"
)

func (h *Handler) createTitle(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorMessageResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input news.Title
	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TitleList.Create(id.(int), input)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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
