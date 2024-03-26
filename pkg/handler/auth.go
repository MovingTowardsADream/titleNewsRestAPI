package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	news "titleNewsRestApi"
)

func (h *Handler) signUp(c *gin.Context) {
	// Getting and validation input
	var input news.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
