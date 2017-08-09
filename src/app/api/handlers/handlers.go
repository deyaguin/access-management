package handlers

import (
	"github.com/labstack/echo"
	"app/models"
	"net/http"
	"app/db"
	"app/api/services"
)

type Handler struct {
	DB db.DB
	Service *services.Service
}

func (h *Handler) CreateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	h.DB.Create(user)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c echo.Context) error {
	users := new([]models.User)
	h.DB.GetAll(users)
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateGroup(c echo.Context) error {
	group := new(models.Group)
	c.Bind(group)
	h.DB.Create(group)
	return c.JSON(http.StatusOK, group)
}

func (h *Handler) GetGroups(c echo.Context) error {
	groups := new([]models.Group)
	h.DB.GetAll(groups)
	return c.JSON(http.StatusOK, groups)
}

func (h *Handler) CreatePolicy(c echo.Context) error {
	policy := new(models.Policy)
	c.Bind(policy)
	h.DB.Create(policy)
	return c.JSON(http.StatusOK, policy)
}

func (h *Handler) GetPolicies(c echo.Context) error {
	policies := new([]models.Policy)
	h.DB.GetAll(policies)
	return c.JSON(http.StatusOK, policies)
}

func (h *Handler) Sm(c echo.Context) error {
	policies := new([]models.Policy)
	users := new([]models.User)
	h.DB.GetAll(users)
	user := (*users)[0]
	h.Service.GetUserPolicies(&user, policies)
	return c.JSON(http.StatusOK, policies)
}
