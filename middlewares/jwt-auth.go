package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"example.com/sagor/go-web-gin/service"
	"example.com/sagor/go-web-gin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.Request.Header.Get("Authorization")
		if strings.TrimSpace(authHeader) == "" {
			utils.ReturnUnauthorized(ctx, http.StatusUnauthorized, "You are not authorized to access this path!!", fmt.Errorf("unauthorized access"))
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			utils.ReturnUnauthorized(ctx, http.StatusUnauthorized, "You are not authorized to access this path!!", err)
			return
		}
	}
}
