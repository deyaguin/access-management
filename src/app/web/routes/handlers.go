package routes

import (
	"github.com/labstack/echo"
	"../models"
)

func sm(c echo.Context) error {
	user := &models.User{
		Name: c.FormValue("name"),
	}
}
