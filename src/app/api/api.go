package api

import (
	"github.com/labstack/echo"
	"app/db"
)

type Api struct {
	DB db.DB
}

func (a *Api) Init() {
	e := echo.New()

	e.POST("/users", a.CreateUser)
	e.GET("/users", a.GetUsers)
	e.POST("/groups", a.CreateGroup)
	e.GET("/groups", a.GetGroups)
	e.POST("/policies", a.CreatePolicy)
	e.GET("/policies", a.GetPolicies)
	e.POST("/sm", a.CheckPermissions)

	e.Start(":1535")
}
