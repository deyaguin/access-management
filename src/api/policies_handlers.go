package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

func (a *Api) createPolicy(c echo.Context) error {
	policyCreating := &models.Policy{}

	if err := c.Bind(policyCreating); err != nil {
		return NewUnprocessableBodyError()
	}

	policy, err := a.policyService.CreatePolicy(policyCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, policy)
}

func (a *Api) getPolicies(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError("page", c.QueryParam("page"))
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError("per_page", c.QueryParam("per_page"))
	}

	policies, err := a.policyService.GetPolicies(page, perPage)
	if err != nil {
		return NewInvalidQueryError("GroupID", string(3))
	}

	return c.JSON(http.StatusOK, policies)
}

func (a *Api) getPolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", string(id))
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
		return NewInvalidQueryError("PolicyID", c.Param("id"))
	}

	policyUpdating := &models.Policy{ID: id}
	if err := c.Bind(policyUpdating); err != nil {
		return NewUnprocessableBodyError()
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
		return NewInvalidQueryError("PolicyID", c.Param("id"))
	}

	if err := a.policyService.RemovePolicy(id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) addPermissionsToPolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", c.Param("id"))
	}
	policy := &models.Policy{ID: id}

	permissions := new(permissions)
	if err = c.Bind(permissions); err != nil {
		return NewUnprocessableBodyError()
	}

	if err = a.policyService.AddPermissionsToPolicy(policy, permissions.Permissions); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) removePermissionFromPolicy(c echo.Context) error {
	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", c.Param("policyId"))
	}

	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		return NewInvalidQueryError("PermissionID", c.Param("permissionId"))
	}

	if err = a.policyService.RemovePermissionFromPolicy(policyId, permissionId); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) getPermissionsByPolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", c.Param("id"))
	}

	permissions, err := a.policyService.GetPermissionsByPolicy(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, permissions)
}
