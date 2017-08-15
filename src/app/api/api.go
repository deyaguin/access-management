package api

import (
	"github.com/labstack/echo"
	"app/db"
	"gopkg.in/go-playground/validator.v9"
)

type (
	Api struct {
		DB db.DB
	}
	MyValidator struct {
		validator *validator.Validate
	}
)

func (cV *MyValidator) Validate(i interface{}) error {
	return cV.validator.Struct(i)
}


func (a *Api) Init() {
	e := echo.New()
	e.Validator = &MyValidator{validator: validator.New()}

	e.POST("/users", a.createUser)
	e.GET("/users", a.getUsers)
	e.DELETE("/users", a.deleteUser)
	e.POST("/groups", a.createGroup)
	e.GET("/groups", a.getGroups)
	e.POST("/policies", a.createPolicy)
	e.GET("/policies", a.getPolicies)
	e.POST("/checkPermissions", a.userPermissions)
	e.Logger.Fatal(e.Start(":1535"))
}
