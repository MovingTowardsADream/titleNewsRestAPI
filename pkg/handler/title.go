package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	news "titleNewsRestApi"
)

func (h *Handler) createTitle(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input news.Title
	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TitleList.Create(userId, input)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []news.Title `json:"data"`
}

func (h *Handler) getAllTitle(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	lists, err := h.service.TitleList.GetAll(userId)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getTitleById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.service.TitleList.GetById(userId, id)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

//func (h *Handler) updateTitle(c *gin.Context) {
//
//}

func (h *Handler) deleteTitle(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.TitleList.Delete(userId, id)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}
