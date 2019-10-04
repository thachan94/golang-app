package userController

import (
	"golang-app/models/userModel"
	"golang-app/utility/hash"

	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var getUser = new(userModel.UserModel)

type UserController struct{}

func (con *UserController) Login(e echo.Context) (err error) {
	user := &userModel.User{}
	if err = e.Bind(user); err != nil {
		return
	}
	requestPassword := user.Password
	err = getUser.Login(user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid username"}
		}
	}
	if !hash.CheckPasswordHash(requestPassword, user.Password) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid password"}
	} else {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
		user.Token, err = token.SignedString([]byte("user-secret"))
		if err != nil {
			return err
		}
		return e.JSON(http.StatusOK, echo.Map{"token": user.Token})
	}
}

func (con *UserController) Signup(e echo.Context) (err error) {
	user := &userModel.User{}
	if err = e.Bind(user); err != nil {
		return
	}
	count, err := getUser.CheckUserExist(user)
	if err != nil {
		fmt.Println(err)
	}
	if count == 0 {
		user.Password, err = hash.HashPassword(user.Password)
		user.CreateTime = time.Now().Unix()
		if err != nil {
			return
		}
		err = getUser.Signup(user)
		if err != nil {
			fmt.Println(err)
		}
		return e.JSON(http.StatusOK, echo.Map{"message": "success"})
	} else {
		return e.JSON(http.StatusOK, echo.Map{"message": "exist"})
	}
}

func (con *UserController) UserInfo(e echo.Context) (err error) {
	user := &userModel.User{}
	if err = e.Bind(user); err != nil {
		return
	}
	user.ID = bson.ObjectIdHex(userIDFromToken(e))

	result, err := getUser.GetUserInfo(user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid username"}
		}
	}
	return e.JSON(http.StatusOK, echo.Map{"username": result.UserName, "password": result.Password})
}

func userIDFromToken(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func (con *UserController) GetID(e echo.Context) (err error) {
	userID := userIDFromToken(e)
	return e.JSON(http.StatusOK, echo.Map{"id": userID})
}
