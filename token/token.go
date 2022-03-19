package token

import (
	"fmt"
	"time"

	dto "example.com/sagor/go-web-gin/dto"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	JWTPrivateKey = "SercretTokenSecretToken"
	ip            = "192.168.0.107"
)

func GenerateToken(claims *dto.JwtClaims, expirationTime time.Time) (string, error) {
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

func VerifyToken(tokenString string) (bool, *dto.JwtClaims) {
	claims := &dto.JwtClaims{}
	token, _ := getTokenFromString(tokenString, claims)
	// fmt.Println(claims)
	if token.Valid {
		if err := claims.Valid(); err == nil {
			return true, claims
		}
	}
	return false, claims
}

func GetClaims(tokenString string) dto.JwtClaims {
	claims := &dto.JwtClaims{}
	getTokenFromString(tokenString, claims)
	return *claims
}

func getTokenFromString(tokenString string, claims *dto.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// validate alg
		// fmt.Println(token.Header["alg"])
		// fmt.Println(token.Method.Alg())
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// return secret key so that the token can be validated by 'ParseWithClaims' method
		return []byte(JWTPrivateKey), nil
	})
}
