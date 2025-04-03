package auth

import (
	"awesomeProject/config"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expriration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expriration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
