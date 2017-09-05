package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

func (a *API) createPolicy(c echo.Context) error {
	policyCreating := &models.Policy{}

	if err := c.Bind(policyCreating); err != nil {
		return NewUnprocessableBodyError()
	}

	policy, err := a.policiesService.CreatePolicy(policyCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, policy)
}

func (a *API) getPolicies(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError(
			"page",
			c.QueryParam("page"),
		)
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError(
			"per_page",
			c.QueryParam("per_page"),
		)
	}

	if err := checkPaginationParams(page, perPage); err != nil {
		return err
	}

	policies, err := a.policiesService.GetPolicies(page, perPage)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}

func (a *API) getPolicy(c echo.Context) error {
	policyID, err := strconv.Atoi(c.Param("policyID"))
	if err != nil {
		return NewInvalidQueryError(
			"PolicyID",
			c.Param("policyID"),
		)
	}

	policy, err := a.policiesService.GetPolicy(policyID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policy)
}

func (a *API) updatePolicy(c echo.Context) error {
	policyID, err := strconv.Atoi(c.Param("policyID"))
	if err != nil {
		return NewInvalidQueryError(
			"PolicyID",
			c.Param("policyID"),
		)
	}

	policyUpdating := &models.Policy{ID: policyID}
	if err := c.Bind(policyUpdating); err != nil {
		return NewUnprocessableBodyError()
	}

	policy, err := a.policiesService.UpdatePolicy(policyUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policy)
}

func (a *API) removePolicy(c echo.Context) error {
	policyID, err := strconv.Atoi(c.Param("policyID"))
	if err != nil {
		return NewInvalidQueryError(
			"PolicyID",
			c.Param("policyID"),
		)
	}

	if err := a.policiesService.RemovePolicy(policyID); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *API) addPermissionsToPolicy(c echo.Context) error {
	policyID, err := strconv.Atoi(c.Param("policyID"))
	if err != nil {
		return NewInvalidQueryError(
			"PolicyID",
			c.Param("policyID"),
		)
	}
	policy := &models.Policy{ID: policyID}

	permissions := new(permissions)
	if err = c.Bind(permissions); err != nil {
		return NewUnprocessableBodyError()
	}

	if err = a.policiesService.AddPermissionsToPolicy(
		policy,
		permissions.Permissions,
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *API) removePermissionFromPolicy(c echo.Context) error {
	policyID, err := strconv.Atoi(c.Param("policyID"))
	if err != nil {
		return NewInvalidQueryError(
			"PolicyID",
			c.Param("policyID"),
		)
	}

	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		return NewInvalidQueryError(
			"PermissionID",
			c.Param("permissionId"),
		)
	}

	if err = a.policiesService.RemovePermissionFromPolicy(
		policyID,
		permissionId,
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *API) getPermissionsByPolicy(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError(
			"page",
			c.QueryParam("page"),
		)
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError(
			"per_page",
			c.QueryParam("per_page"),
		)
	}

	if err := checkPaginationParams(page, perPage); err != nil {
		return err
	}

	policyID, err := strconv.Atoi(c.Param("policyID"))
	if err != nil {
		return NewInvalidQueryError(
			"PolicyID",
			c.Param("policyID"),
		)
	}

	permissions, err := a.policiesService.GetPermissionsByPolicy(
		policyID,
		page,
		perPage,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, permissions)
}
