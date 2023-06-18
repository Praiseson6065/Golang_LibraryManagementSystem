package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"golang.org/x/crypto/bcrypt"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	jtoken "github.com/golang-jwt/jwt/v4"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func CookieGetData(cookie string, c *fiber.Ctx) (jtoken.MapClaims, error) {

	token, err := jtoken.Parse(cookie, func(token *jtoken.Token) (interface{}, error) {

		return []byte(config.Secret), nil
	})
	if err != nil {

		return jtoken.MapClaims{
				"name": "nologin",
			}, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
	}

	claims, ok := token.Claims.(jtoken.MapClaims)
	if !ok || !token.Valid {

		return jtoken.MapClaims{
				"name": "nologin",
			}, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
	}

	return claims, err
}
