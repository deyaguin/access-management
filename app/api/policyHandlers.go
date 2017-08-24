package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
	"strconv"
)

type policyParams struct {
	Name string `validation:"required"`
}

func (a *Api) createPolicy(c echo.Context) error {
	pParams := new(policyParams)
	c.Bind(pParams)
	if err := c.Validate(pParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	policy := &models.Policy{Name: pParams.Name}
	if err := a.DB.CreatePolicy(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, policy)
}

func (a *Api) getPolicies(c echo.Context) error {
	policies, err := a.DB.GetPolicies()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, policies)
}

func (a *Api) getPolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	policy, err := a.DB.GetPolicy(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) updatePolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err := a.DB.GetPolicy(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	pParams := new(policyParams)
	c.Bind(pParams)
	if err = c.Validate(pParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var policy = &models.Policy{ID: id, Name: pParams.Name}
	if err = a.DB.UpdatePolicy(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) removePolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if _, err := a.DB.GetPolicy(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	policy := &models.Policy{ID: id}
	if err = a.DB.DeletePolicy(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}
