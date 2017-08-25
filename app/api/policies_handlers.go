package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
	"strconv"
)

func (a *Api) createPolicy(c echo.Context) error {
	policyCreating := &models.Policy{}

	if err := c.Bind(policyCreating); err != nil {
		return err
	}

	policy, err := a.policyService.CreatePolicy(policyCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, policy)
}

func (a *Api) getPolicies(c echo.Context) error {
	policies, err := a.policyService.GetPolicies()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}

func (a *Api) getPolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	policy, err := a.policyService.GetPolicy(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policy)
}

func (a *Api) updatePolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	policyUpdating := &models.Policy{ID: id}
	if err := c.Bind(policyUpdating); err != nil {
		return err
	}

	policy, err := a.policyService.UpdatePolicy(policyUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policy)
}

func (a *Api) removePolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	policy := &models.Policy{ID: id}
	if err := a.policyService.RemovePolicy(policy); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
