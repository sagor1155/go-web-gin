package controller

import (
	"fmt"
	"net/http"

	"example.com/sagor/go-web-gin/dto"
	"example.com/sagor/go-web-gin/service"
	utils "example.com/sagor/go-web-gin/utils"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) {
	credentials := dto.Credentials{}
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		utils.ReturnBadRequest(ctx, http.StatusBadRequest, "Invalid user Credentials", err)
		return
	}

	if isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password); !isAuthenticated {
		utils.ReturnBadRequest(ctx, http.StatusBadRequest, "Unauthenticated user!!", fmt.Errorf("unauthenticated user"))
		return
	}
	token := controller.jwtService.GenerateToken(credentials.Username, true)
	utils.ReturnOK(ctx, http.StatusOK, "Login Successfull", token)
}
