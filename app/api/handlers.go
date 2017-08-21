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

func (a *Api) userPermissions(c echo.Context) (err error) {
	cP := new(checkParams)
	c.Bind(cP)
	access, err := a.check(cP)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, access)
}
