package api

import (
	"github.com/labstack/echo"
	"app/api/handlers"
	"app/db"
	"app/api/services"
)

type Api struct {
	DB db.DB
}

func (api *Api) Init() {
	e := echo.New()
	s := services.Service{
		api.DB,
	}
	h := handlers.Handler{
		api.DB,
		&s,
	}

	e.POST("/users", h.CreateUser)
	e.GET("/users", h.GetUsers)
	e.POST("/groups", h.CreateGroup)
	e.GET("/groups", h.GetGroups)
	e.POST("/policies", h.CreatePolicy)
	e.GET("/policies", h.GetPolicies)
	e.GET("/sm", h.Sm)

	e.Start(":1535")
}
