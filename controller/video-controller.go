package controller

import (
	"net/http"

	"example.com/sagor/go-web-gin/entity"
	"example.com/sagor/go-web-gin/service"
	utils "example.com/sagor/go-web-gin/utils"
	"example.com/sagor/go-web-gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	utils.OK(ctx, http.StatusOK, "Successfull", c.service.FindAll())
}

func (c *controller) Save(ctx *gin.Context) {
	var video entity.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		utils.BadRequest(ctx, http.StatusBadRequest, "Invalid video format!!", err)
		return
	}

	if err := validate.Struct(video); err != nil {
		utils.BadRequest(ctx, http.StatusBadRequest, "Invalid video format!!", err)
		return
	}

	utils.OK(ctx, http.StatusOK, "Video saved successfully!!", c.service.Save(video))
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
