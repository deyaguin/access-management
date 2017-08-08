package api

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"app/api/handlers"
)

type Api struct {
	DB *gorm.DB
}

func (api *Api) Init() {
	e := echo.New()
	h := &handlers.Handler{
		api.DB,
	}
	e.POST("/users", h.CreateUser)
	e.GET("/users", h.GetUsers)

	e.Start(":1535")
}
