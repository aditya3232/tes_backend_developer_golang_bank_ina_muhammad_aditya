package middleware

import (
	"context"
	"net/http"
	"strings"

	oauth2_library "github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/library/oauth2"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func Oauth2Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Split the Authorization header to get the token
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Validate the token
		token := splitToken[1]
		googleConfig := oauth2_library.GoogleConfig()
		client := googleConfig.Client(context.Background(), &oauth2.Token{AccessToken: token})
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Add user information to the context for further use in the handler
		c.Set("user", resp)

		// Continue to the next middleware or the actual handler
		c.Next()
	}
}
