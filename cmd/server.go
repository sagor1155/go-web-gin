package main

import (
	"io"
	"net/http"
	"os"

	"example.com/sagor/go-web-gin/controller"
	"example.com/sagor/go-web-gin/middlewares"
	"example.com/sagor/go-web-gin/service"
	"github.com/gin-gonic/gin"
	// gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

const (
	ROLE_ADMIN = 1
	ROLE_USER  = 2
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // write log output into both file and stdout
}

func main() {
	setupLogOutput()

	server := gin.New()

	// load static assests
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	// add middleware
	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	server.Use(gin.Recovery(), middlewares.Logger())

	server.POST("/login", controller.Login)

	apiRoutes := server.Group("/api")
	apiRoutes.Use(middlewares.ValidateToken())
	// apiRoutes.Use(middlewares.Authorization([]int{ROLE_ADMIN, ROLE_USER}))

	apiRoutes.GET("/videos", middlewares.Authorization([]int{ROLE_USER}), func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	apiRoutes.POST("/videos", middlewares.Authorization([]int{ROLE_ADMIN}), func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, gin.H{"message": "Video saved successfully!!"})
		}
	})

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
