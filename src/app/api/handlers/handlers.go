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
	user := new(models.User)
	c.Bind(user)
	h.DB.CreateUser(user)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c echo.Context) error {
	users := new([]models.User)
	h.DB.GetUsers(users)
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateGroup(c echo.Context) error {
	group := new(models.Group)
	c.Bind(group)
	h.DB.CreateGroup(group)
	return c.JSON(http.StatusOK, group)
}

func (h *Handler) GetGroups(c echo.Context) error {
	groups := new([]models.Group)
	h.DB.GetGroups(groups)
	return c.JSON(http.StatusOK, groups)
}

func (h *Handler) CreatePolicy(c echo.Context) error {
	policy := new(models.Policy)
	c.Bind(policy)
	h.DB.CreatePolicy(policy)
	return c.JSON(http.StatusOK, policy)
}

func (h *Handler) GetPolicies(c echo.Context) error {
	policies := new([]models.Policy)
	h.DB.GetPolicies(policies)
	return c.JSON(http.StatusOK, policies)
}
