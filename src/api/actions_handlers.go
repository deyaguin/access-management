package api

import (
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (a *API) createAction(c echo.Context) error {
	actionCreating := &models.Action{}

	if err := c.Bind(actionCreating); err != nil {
		return NewUnprocessableBodyError()
	}

	action, err := a.CreateAction(actionCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, action)
}

func (a *API) getActions(c echo.Context) error {
	actions, err := a.GetActions()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, actions)
}

func (a *API) getAction(c echo.Context) error {
	actionId, err := strconv.Atoi(c.Param("actionID"))
	if err != nil {
		return NewInvalidQueryError(
			"actionID",
			c.Param("actionID"),
		)
	}

	action, err := a.GetAction(actionId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, action)
}

func (a *API) updateAction(c echo.Context) error {
	actionID, err := strconv.Atoi(c.Param("actionID"))
	if err != nil {
		return NewInvalidQueryError(
			"actionID",
			c.Param("actionID"),
		)
	}

	actionUpdating := &models.Action{ID: actionID}
	if err := c.Bind(actionUpdating); err != nil {
		return NewUnprocessableBodyError()
	}

	group, err := a.UpdateAction(actionUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, group)
}

func (a *API) removeAction(c echo.Context) error {
	actionID, err := strconv.Atoi(c.Param("actionID"))
	if err != nil {
		return NewInvalidQueryError(
			"actionID",
			c.Param("actionID"),
		)
	}

	if err := a.RemoveAction(actionID); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
