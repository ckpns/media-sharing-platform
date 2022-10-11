package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/ckpns/media-sharing-platform/backend/src/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthRequired() gin.HandlerFunc {

	return func(context *gin.Context) {

		token := extractToken(context)
		userID, err := extractIDFromToken(token)
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		context.Set("user_id", userID.String())
	}
}

func extractToken(context *gin.Context) string {

	bearerToken := context.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func extractIDFromToken(token string) (uuid.UUID, error) {

	config := common.GetConfig()

	extractedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.AUTH.HmacSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := extractedToken.Claims.(jwt.MapClaims)
	if !ok || !extractedToken.Valid {
		return uuid.UUID{}, errors.New("error")
	}

	id, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}
