package auth

import (
	"awesomeProject/config"
	"awesomeProject/types"
	"awesomeProject/utils"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type contextKey string

const UserKey contextKey = "userID"

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

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get the token from the user request
		tokenString := getTokenFromRequest(r)
		//validate the jwt
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}
		if !token.Valid {
			log.Printf("invalid token")
			permissionDenied(w)
			return
		}
		//if it is we fetch the userID from the DB(id from the token)
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)
		userID, _ := strconv.Atoi(str)

		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user: %v", err)
			permissionDenied(w)
			return
		}

		//set context "userID" to the user ID
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth != "" {
		print(tokenAuth)
		return tokenAuth
	}
	return ""
}
func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userID
}
