package token

import (
	"fmt"
	"time"

	"example.com/sagor/go-web-gin/entity"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	JWTPrivateKey = "SercretTokenSecretToken"
	ip            = "192.168.0.107"
)

func GenerateToken(claims *entity.JwtClaims, expirationTime time.Time) (string, error) {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWTPrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, *entity.JwtClaims) {
	claims := &entity.JwtClaims{}
	token, _ := getTokenFromString(tokenString, claims)
	// fmt.Println(claims)
	if token.Valid {
		if err := claims.Valid(); err == nil {
			return true, claims
		}
	}
	return false, claims
}

func GetClaims(tokenString string) entity.JwtClaims {
	claims := &entity.JwtClaims{}
	getTokenFromString(tokenString, claims)
	return *claims
}

func getTokenFromString(tokenString string, claims *entity.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// validate alg
		// fmt.Println(token.Header["alg"])
		// fmt.Println(token.Method.Alg())
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTPrivateKey), nil
	})
}
