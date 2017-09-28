package api

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (a *API) getAllServices(c echo.Context) error {
	services, err := a.GetAllServices()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, services)
}

func (a *API) getActionsByService(c echo.Context) error {
	serviceID, err := strconv.Atoi(c.Param("serviceID"))
 	if err != nil {
		return NewInvalidQueryError(
			"ServiceID",
			c.Param("serviceID"),
		)
 	}

 	actions, err := a.GetActionsByService(serviceID)
 	if err != nil {
 		return err
	}

	return c.JSON(http.StatusOK, actions)
}
