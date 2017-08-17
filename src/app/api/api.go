package api

import (
	"github.com/labstack/echo"
	"app/db"
	"gopkg.in/go-playground/validator.v9"
)

type (
	Api struct {
		DB db.DB
	}
	MyValidator struct {
		validator *validator.Validate
	}
)

func (cV *MyValidator) Validate(i interface{}) error {
	return cV.validator.Struct(i)
}

func (a *Api) Init() {
	e := echo.New()
	e.Validator = &MyValidator{validator: validator.New()}

	e.POST("/users", a.createUser)
	e.GET("/users", a.getUsers)
	e.PUT("/users", a.updateUser)
	e.DELETE("/users", a.deleteUser)
	e.POST("/groups", a.createGroup)
	e.GET("/groups", a.getGroups)
	e.PUT("/groups", a.updateGroup)
	e.DELETE("groups", a.deleteGroup)
	e.POST("/policies", a.createPolicy)
	e.GET("/policies", a.getPolicies)
	e.PUT("/policies", a.updatePolicy)
	e.DELETE("/policies", a.deletePolicy)

	e.POST("/addUserToGroups", a.addUserToGroups)
	e.POST("/addUsersToGroup", a.addUsersToGroup)
	e.POST("/attachPoliciesByUser", a.attachPoliciesByUser)
	e.POST("/attachPoliciesByGroup", a.attachPoliciesByGroup)
	e.POST("/attachPolicyToUsers", a.attachPolicyToUsers)
	e.POST("/attachPolicyToGroups", a.attachPolicyToGroups)
	e.POST("/removeUserFromGroups", a.removeUserFromGroups)
	e.POST("/removeUsersFromGroup", a.removeUsersFromGroup)
	e.POST("/detachPoliciesByUser", a.detachPoliciesByUser)
	e.POST("/detachPoliciesByGroup", a.detachPoliciesByGroup)
	e.POST("/detachUsersFromPolicy", a.detachUsersFromPolicy)
	e.POST("/detachGroupsFromPolicy", a.detachGroupsFromPolicy)

	e.POST("/checkPermissions", a.userPermissions)
	e.Logger.Fatal(e.Start(":1535"))
}
