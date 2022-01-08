package utils

import (
	"errors"
	"fmt"
	"majoo-backend-test/app/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Generate JWT Token
func GenerateToken(user models.UserResponse) (string, error) {
	// Get JWT Key From ENV
	key := os.Getenv("JWT_SECRET_KEY")

	// Create New Claims
	claims := jwt.MapClaims{}

	// Assign data to Claims
	claims["authorized"] = true
	claims["user_id"] = user.Id
	claims["user_name"] = user.UserName
	claims["expired"] = time.Now().Add(time.Hour * 24).Unix() // Valid for 24 Hours

	// Prepare JWT
	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create JWT ( Token )
	token, err := unsignedJWT.SignedString([]byte(key))

	// Check if there is error when signing JWT with Secret Key
	if err != nil {
		return "", err
	}

	// Return JWT ( Token )
	return token, nil
}

// Extract Authorization Token
func ExtractToken(context *gin.Context) string {
	// Get Token from Header
	auhorizationToken := context.Request.Header.Get("Authorization")

	// Return Token
	return auhorizationToken
}

// Get User Detail From Token
func GetUserDetailFromToken(context *gin.Context) (models.UserClaims, error) {
	// Extract Token Data
	tokenString := ExtractToken(context)

	// Get Secret Key from ENV
	key := os.Getenv("JWT_SECRET_KEY")

	// User Claims Object
	var user models.UserClaims

	// Parse JWT and validate
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signin method token
		if jwt.GetSigningMethod("HS256") != token.Method {
			// When signin method not same
			return nil, fmt.Errorf("jwt: unexpected signing method, %v", token.Header["alg"])
		}

		// return key
		return []byte(key), nil
	})

	// Check if user exist in database & Token Expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Catch User Claims from token
		userId := int(claims["user_id"].(float64))
		userName := fmt.Sprintf("%v", claims["user_name"])

		// Assign to User Claims Object
		user.Id = userId
		user.UserName = userName

		// Return Value
		return user, nil
	} else {
		// Return Error
		return user, errors.New("jwt: error parsing jwt")
	}
}
