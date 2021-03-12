package main

import (
	"echo-rest-api/config"
	"echo-rest-api/controllers"
	"echo-rest-api/storage"
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Getting port from CLI
	url := config.GetDefaultApiUrl()

	// Printing out the configuration
	log.Info("* API URL: " + url)
	log.Info("* JWT Secret: " + config.GetDefaultJWTSecret())
	log.Infof("* Rate limiting: %f req/s", config.GetDefaultRequestsPerSecond())
	log.Info("* Database configuration: " + config.GetPostgresConnectionString())

	// Initializing echo
	e := echo.New()

	// Initializing database
	storage.InitDB()

	// Logging
	e.Use(middleware.Logger())

	// Recover from errors
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// CSRF (using default configuration)
	e.Use(middleware.CSRF())

	// Rate limiting
	limiter := tollbooth.NewLimiter(config.GetDefaultRequestsPerSecond(), nil) // 1 request per second
	limiter.SetMessage("{\"message\":\"You have reached maximum request limit\"}\n")
	limiter.SetMessageContentType("application/json; charset=utf-8")
	limiter.SetMethods([]string{"GET", "POST", "PUT", "DELETE"})

	// It works!
	e.GET("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!\n")
	}), tollbooth_echo.LimitHandler(limiter))

	// Unauthenticated routes (CRUD)
	e.GET("/users", echo.HandlerFunc(controllers.GetUsers), tollbooth_echo.LimitHandler(limiter))
	e.GET("/users/:id", echo.HandlerFunc(controllers.GetUser), tollbooth_echo.LimitHandler(limiter))
	e.POST("/users", echo.HandlerFunc(controllers.CreateUser), tollbooth_echo.LimitHandler(limiter))
	e.PUT("/users/:id", echo.HandlerFunc(controllers.UpdateUser), tollbooth_echo.LimitHandler(limiter))
	e.DELETE("/users/:id", echo.HandlerFunc(controllers.DeleteUser), tollbooth_echo.LimitHandler(limiter))

	// Unauthenticated routes (login)
	e.POST("/login", echo.HandlerFunc(controllers.Login), tollbooth_echo.LimitHandler(limiter))

	// Restricted routes (no rate limiting)
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(config.GetDefaultJWTSecret())))
	r.GET("", controllers.RestrictedRoute)

	// Server
	e.Logger.Fatal(e.Start(url))
}
