package middlewares

import (
	"fmt"
	"majoo-backend-test/app/common"
	"majoo-backend-test/app/repositories"
	"majoo-backend-test/app/utils"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Token Authentication
func TokenAuthMiddleware(repository repositories.UserRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Extract Token Data
		tokenString, err := utils.ExtractToken(context)

		// Check if there is error when get user detail
		if err != nil {
			common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})

			// Abort to do next handler
			context.Abort()
			return
		}

		// Check token string
		if tokenString == "" {
			// Return Error
			common.SendError(context, http.StatusUnauthorized, 1, "Unauthorized", []string{"Authentication Token Required"})

			// Abort to do next handler
			context.Abort()
			return
		}

		// Get Secret Key from ENV
		key := os.Getenv("JWT_SECRET_KEY")

		// Parse JWT and validate
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check signin method token
			if jwt.GetSigningMethod("HS256") != token.Method {
				// When signin method not same
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// return key
			return []byte(key), nil
		})

		// Check if user exist in database & Token Expired
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Catch User Email
			userName := fmt.Sprintf("%v", claims["user_name"])

			// Check if token expired
			if time.Now().Unix() > int64(math.Round(claims["expired"].(float64))) {
				// Return Error
				common.SendError(context, http.StatusUnauthorized, 1, "Unauthorized", []string{"Token expired"})

				// Abort to do next handler
				context.Abort()
				return
			}

			// Check User On Database
			userExist, err := repository.CheckIfUserExistsByUsername(userName)

			if !userExist {
				// If Query Error
				if err != nil {
					// Return Error
					common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{err.Error()})

					// Abort to do next handler
					context.Abort()
					return
				}
				// Return Error
				common.SendError(context, http.StatusUnauthorized, 1, "Unauthorized", []string{"User not found"})

				// Abort to do next handler
				context.Abort()
				return
			}

		} else {
			// Token not valid
			common.SendError(context, http.StatusUnauthorized, 1, "Unauthorized", []string{err.Error()})

			// Abort to do next handler
			context.Abort()
			return
		}

		// Check error
		if token == nil && err != nil {
			// Token not valid
			common.SendError(context, http.StatusUnauthorized, 1, "Unauthorized", []string{err.Error()})

			// Abort to do next handler
			context.Abort()
			return
		}

		// Next
		context.Next()
	}
}
