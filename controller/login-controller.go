package controller

import (
	"fmt"
	"net/http"
	"time"

	"example.com/sagor/go-web-gin/entity"
	"example.com/sagor/go-web-gin/token"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var loginObj entity.LoginRequest
	if err := ctx.ShouldBindJSON(&loginObj); err != nil {
		BadRequest(ctx, http.StatusBadRequest, "Invalid Request", err)
		return
	}

	fmt.Println(loginObj)
	// validate the login object for valid credentials and if these are valid then

	claims := &entity.JwtClaims{}
	claims.CompanyId = "CompanyId"
	claims.Username = loginObj.Username
	claims.Roles = []int{1, 2, 3}
	claims.Audience = ctx.Request.Header.Get("Referer")

	tokeCreationTime := time.Now().UTC()
	expirationTime := tokeCreationTime.Add(time.Duration(1) * time.Minute)
	tokenString, err := token.GenerateToken(claims, expirationTime)

	if err != nil {
		BadRequest(ctx, http.StatusBadRequest, "error in generating token", err)
		return
	}
	OK(ctx, http.StatusOK, "token created", tokenString)
}
