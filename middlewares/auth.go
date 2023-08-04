package middlewares

import (
	"fmt"
	"strconv"
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

		return []byte(config.EnvConfigs.SecretKey), nil
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
func UserMiddleWare(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, _ := CookieGetData(cookie, c)
	if claims["usertype"] != "user" {
		return c.SendString("Un Authorized")
	}
	userID := claims["ID"]
	Uid ,err:=strconv.Atoi(c.Params("userid"))
	if err!=nil{
		return c.JSON("Unauthorized")
	}
	s := fmt.Sprintf("%.0f",userID)
	k,err:= strconv.Atoi(s)
	if err!=nil{
		return c.JSON("Unauthorized")
	}
	if  Uid==k {
		return c.Next()
	}

	return c.SendString("Unauthorized")
}
func AdminMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, _ := CookieGetData(cookie, c)
	if claims["usertype"] != "admin" {
		return c.SendString("Un Authorized")
	}
	userID := claims["ID"]
	Uid ,err:=strconv.Atoi(c.Params("userid"))
	if err!=nil{
		return c.JSON("Unauthorized")
	}
	s := fmt.Sprintf("%.0f",userID)
	k,err:= strconv.Atoi(s)
	if err!=nil{
		return c.JSON("Unauthorized")
	}
	if  Uid==k {
		return c.Next()
	}

	return c.JSON("Unauthorized")
}