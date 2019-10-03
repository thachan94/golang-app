package main

import (
	"app/controllers/userController"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	conn := echo.New()

	user := conn.Group("/user")
	{

		newUser := new(userController.UserController)
		user.POST("/login", newUser.Login)
		user.POST("/signup", newUser.Signup)
	}
	info := conn.Group("/info")
	{

		newUser := new(userController.UserController)
		//info.Use(middleware.JWT([]byte("user-secret")))
		//info.Use(middleware.JWT([]byte("user-secret")))
		info.Use(middleware.JWT([]byte("user-secret")))
		info.GET("/id", newUser.GetID)
		info.GET("/user", newUser.UserInfo)
	}
	conn.Logger.Fatal(conn.Start(":8766"))
}
