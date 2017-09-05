package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

func (a *API) createAction(c echo.Context) error {
	actionCreating := &models.Action{}

	if err := c.Bind(actionCreating); err != nil {
		return NewUnprocessableBodyError()
	}

	action, err := a.actionsService.CreateAction(actionCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, action)
}

func (a *API) getActions(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError(
			"page",
			c.QueryParam("page"),
		)
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError("per_page",
			c.QueryParam("per_page"),
		)
	}

	if err := checkPaginationParams(page, perPage); err != nil {
		return err
	}

	//actions, err := a.actionsService.GetActions(page, perPage)
	//if err != nil {
	//
	//}
}
