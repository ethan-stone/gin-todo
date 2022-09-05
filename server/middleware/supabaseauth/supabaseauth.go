package supabaseauth

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

type Claims struct {
	jwt.StandardClaims
	Email string `json:"email,omitempty"`
}

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			log.Error().Msg("No Authorization header")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "No Authorization header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Must provide a bearer token",
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Error().Msg("Invalid signing method")
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			jwtSecret := []byte(os.Getenv("SUPABASE_JWT_SECRET"))

			return jwtSecret, nil
		})

		if err != nil {
			log.Error().Msgf("Authentication failed with error: %v", err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		claims, ok := token.Claims.(*Claims)

		if !ok {
			log.Error().Msg("Invalid request claims")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		log.Info().Interface("claims", claims).Msg("Authentication Successful")

		c.Locals("Claims", claims)

		return c.Next()
	}
}
