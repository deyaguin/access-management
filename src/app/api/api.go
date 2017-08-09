package api

import (
	"github.com/labstack/echo"
	"app/api/handlers"
	"app/db"
)

type Api struct {
	DB db.DB
}

func (api *Api) Init() {
	e := echo.New()
	h := handlers.Handler{
		api.DB,
	}

	e.POST("/users", h.CreateUser)
	e.GET("/users", h.GetUsers)
	e.POST("/groups", h.CreateGroup)
	e.GET("/groups", h.GetGroups)
	e.POST("/policies", h.CreatePolicy)
	e.GET("/policies", h.GetPolicies)
	e.GET("/sm", h.FormListOfPermission)

	e.Start(":1535")
}
