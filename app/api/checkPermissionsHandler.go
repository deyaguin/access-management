package api

import (
	"github.com/labstack/echo"
	"net/http"
)

type checkParams struct {
	Resourse string `validation:"required"`
	Action   int    `validation:"required"`
	ID       int    `validation:"required"`
}

func (a *Api) userPermissions(c echo.Context) error {
	cParams := new(checkParams)
	c.Bind(cParams)
	access, err := a.checkPermissions(cParams)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, access)
}
