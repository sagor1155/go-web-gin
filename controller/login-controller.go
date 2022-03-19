package controller

import (
	"fmt"
	"net/http"
	"time"

	dto "example.com/sagor/go-web-gin/dto"
	"example.com/sagor/go-web-gin/token"
	utils "example.com/sagor/go-web-gin/utils"
	"github.com/gin-gonic/gin"
)

const (
	username = "DEVSoL"
	password = "SoLDEV"
)

func Login(ctx *gin.Context) {
	var loginObj dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginObj); err != nil {
		utils.BadRequest(ctx, http.StatusBadRequest, "Invalid Request", err)
		return
	}

	// Validate the login object for valid credentials
	if loginObj.Username != username || loginObj.Password != password {
		utils.BadRequest(ctx, http.StatusUnauthorized, "Failed to login!!", fmt.Errorf("invalid user credential"))
		return
	}

	claims := &dto.JwtClaims{}
	claims.Username = loginObj.Username
	claims.Roles = loginObj.Roles
	claims.Audience = ctx.Request.Header.Get("Referer")

	tokenCreationTime := time.Now().UTC()
	expirationTime := tokenCreationTime.Add(time.Duration(1) * time.Minute)
	tokenString, err := token.GenerateToken(claims, expirationTime)

	if err != nil {
		utils.BadRequest(ctx, http.StatusBadRequest, "Error in generating token", err)
		return
	}
	utils.OK(ctx, http.StatusOK, "token created", tokenString)
}
