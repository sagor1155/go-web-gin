package controller

import (
	"net/http"

	"example.com/sagor/go-web-gin/entity"
	"github.com/gin-gonic/gin"
)

func OK(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, entity.Response{
		Data:    data,
		Status:  status,
		Message: message,
		Error:   "null",
	})
}

func BadRequest(ctx *gin.Context, status int, message string, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.Response{
		Error:   err.Error(),
		Status:  status,
		Message: message,
		Data:    "null",
	})
}
