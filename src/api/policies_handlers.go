package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

type permissions struct {
	Permissions *[]models.Permission `validate:"required"`
}

func (a *Api) createPolicy(c echo.Context) error {
	policyCreating := &models.Policy{}

	if err := c.Bind(policyCreating); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
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
		return NewInvalidQueryError("PolicyID", string(id))
	}

	policyUpdating := &models.Policy{ID: id}
	if err := c.Bind(policyUpdating); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
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
		return NewInvalidQueryError("PolicyID", string(id))
	}

	policy := &models.Policy{ID: id}
	if err := a.policyService.RemovePolicy(policy); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) addPermissionsToPolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", string(id))
	}
	policy := &models.Policy{ID: id}

	permissions := new(permissions)
	if err = c.Bind(permissions); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
	}

	if err = a.policyService.AddPermissionsToPolicy(policy, permissions.Permissions); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) removePermissionFromPolicy(c echo.Context) error {
	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", string(policyId))
	}
	policy := &models.Policy{ID: policyId}

	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		return NewInvalidQueryError("PermissionID", string(permissionId))
	}
	permission := &models.Permission{ID: permissionId}

	if err = a.policyService.RemovePermissionFromPolicy(policy, permission); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) getPermissionsByPolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", string(id))
	}
	policy := &models.Policy{ID: id}

	permissions, err := a.policyService.GetPermissionsByPolicy(policy)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, permissions)
}
