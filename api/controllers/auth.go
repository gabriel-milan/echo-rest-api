package controllers

import (
	"echo-rest-api/config"
	"echo-rest-api/lib"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func IsAuthorized(username string, password string) bool {
	// Authorization logic here
	if username == "test" && password == "test" {
		return true
	}
	return false
}

func Login(c echo.Context) error {
	// Parse input data
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Check for authorization
	if !IsAuthorized(username, password) {
		return echo.ErrUnauthorized
	}

	// Generate new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Test user"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(config.GetDefaultJWTSecret()))
	if err != nil {
		errorMessage := lib.Message{Message: "Failed to generate JWT: " + err.Error()}
		return c.JSON(http.StatusInternalServerError, errorMessage)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func RestrictedRoute(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!\n")
}
