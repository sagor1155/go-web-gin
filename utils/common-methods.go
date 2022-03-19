package controller

import (
	"net/http"

	"example.com/sagor/go-web-gin/dto"
	"github.com/gin-gonic/gin"
)

func OK(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, dto.Response{
		Data:    data,
		Status:  status,
		Message: message,
		Error:   "null",
	})
}

func BadRequest(ctx *gin.Context, status int, message string, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
		Error:   err.Error(),
		Status:  status,
		Message: message,
		Data:    "null",
	})
}
