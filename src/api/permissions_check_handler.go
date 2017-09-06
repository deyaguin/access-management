package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/services"
	"net/http"
)

func (a *API) userPermissions(c echo.Context) error {
	checkingParams := new(services.CheckingParams)
	if err := c.Bind(checkingParams); err != nil {
		return NewUnprocessableBodyError()
	}

	access, err := a.CheckPermissions(checkingParams)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		Access bool `json:"access"`
	}{access})
}
