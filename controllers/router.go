package controllers

import (
	"github.com/labstack/echo"
)

func Setup(router *echo.Router) {
	users := UsersController{Router: router}
	users.Setup()
}
