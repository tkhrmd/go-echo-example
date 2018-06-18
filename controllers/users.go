package controllers

import (
	"github.com/labstack/echo"
)

type UsersController struct {
	Router *echo.Router
}

func (controller *UsersController) createUser(c echo.Context) error {
	return nil
}

func (controller *UsersController) Setup() {
	controller.Router.Add("POST", "/users", controller.createUser)
}
