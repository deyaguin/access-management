package handlers

import (
	"github.com/labstack/echo"
	"app/models"
	"net/http"
	"app/db"
)

type Handler struct {
	DB db.DB
}

func (h *Handler) CreateUser(c echo.Context) error {
	user := &models.User{
		Name: c.FormValue("name"),
	}
	h.DB.CreateUser(user)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c echo.Context) error {
	users := &[]models.User{}
	h.DB.GetUsers(users)
	return c.JSON(http.StatusOK, users)
}
