package main

import (
	"golang-app/controllers/userController"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	conn := echo.New()
	conn.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "method=${method}, uri=${uri}, status=${status}\n"}))
	conn.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	user := conn.Group("/user")
	{

		newUser := new(userController.UserController)
		user.POST("/login", newUser.Login)
		user.POST("/signup", newUser.Signup)
	}
	info := conn.Group("/info")
	{

		newUser := new(userController.UserController)
		info.Use(middleware.JWT([]byte("user-secret")))
		info.GET("/id", newUser.GetID)
		info.GET("/user", newUser.UserInfo)
	}
	conn.Logger.Fatal(conn.Start(":8766"))
}
