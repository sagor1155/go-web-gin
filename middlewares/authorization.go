package middlewares

import (
	"net/http"

	"example.com/sagor/go-web-gin/entity"
	"example.com/sagor/go-web-gin/token"
	"github.com/gin-gonic/gin"
)

func ReturnUnauthorized(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, entity.Response{
		Error:   "You are not authorized to access this path",
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
		Data:    "null",
	})
}

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("apikey")
		// referer := ctx.Request.Header.Get("Referer")
		valid, claims := token.VerifyToken(tokenString)
		if !valid {
			ReturnUnauthorized(ctx)
		} else {
			if len(ctx.Keys) == 0 {
				ctx.Keys = make(map[string]interface{})
			}
			ctx.Keys["CompanyId"] = claims.CompanyId
			ctx.Keys["Username"] = claims.Username
			ctx.Keys["Roles"] = claims.Roles
		}
	}
}
