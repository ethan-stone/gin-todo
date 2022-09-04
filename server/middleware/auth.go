package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

type Claims struct {
		jwt.StandardClaims
		Email string `json:"email,omitempty"`
}

func SupabaseAuth() gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			log.Error("No Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Must provide a bearer token"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Error("Invalid signing method")
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			jwtSecret := []byte(os.Getenv("SUPABASE_JWT_SECRET"))

			return jwtSecret, nil
		})

		if err != nil {
			log.Errorf("Authentication failed with error: %v", err.Error)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(*Claims)

		if !ok {
			log.Error("Invalid request claims")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		log.WithFields(log.Fields{
			"claims": claims,
		}).Info("Request Authenticated")

		c.Set("Claims", claims)

		c.Next()
	}
}