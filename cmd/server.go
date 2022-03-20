package main

import (
	"io"
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
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
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

	server.POST("/login", loginController.Login)

	apiRoutes := server.Group("/api")
	apiRoutes.Use(middlewares.AuthorizeJWT())
	apiRoutes.GET("/videos", videoController.FindAll)
	apiRoutes.POST("/videos", videoController.Save)

	viewRoutes := server.Group("/view")
	viewRoutes.GET("/videos", videoController.ShowAll)

	server.Run(":8080")
}
