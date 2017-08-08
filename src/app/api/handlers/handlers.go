package handlers

import (
	"github.com/labstack/echo"
	"app/models"
	"net/http"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) CreateUser(c echo.Context) error {
	user := &models.User{
		Name: c.FormValue("name"),
	}
	userDAO := &models.UserDAO{
		h.DB,
	}
	userDAO.Create(user)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c echo.Context) error {
	users := &[]models.User{}
	userDAO := &models.UserDAO{
		h.DB,
	}
	userDAO.Get(users)
	return c.JSON(http.StatusOK, users)
}


//func CreateUser(c echo.Context) error {
//	user := &models.User{
//		Name: c.FormValue("name"),
//	}
//	//userDAO := &models.UserDAO{}
//	return c.JSON(http.StatusOK, user)
//}
