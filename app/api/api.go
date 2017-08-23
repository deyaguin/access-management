package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/db"
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
	e.GET("/users/:id", a.getUser)
	e.PATCH("/users/:id", a.updateUser)
	e.DELETE("/users/:id", a.removeUser)
	e.POST("/groups", a.createGroup)
	e.GET("/groups", a.getGroups)
	e.GET("/groups/:id", a.getGroup)
	e.PATCH("/groups/:id", a.updateGroup)
	e.DELETE("/groups/:id", a.removeGroup)
	e.POST("/policies", a.createPolicy)
	e.GET("/policies", a.getPolicies)
	e.GET("/policies/:id", a.getPolicy)
	e.PATCH("/policies/:id", a.updatePolicy)
	e.DELETE("/policies/:id", a.removePolicy)

	e.PUT("/groups/:id/users", a.addUsersToGroup)
	e.GET("/groups/:id/users", a.getUsersByGroupHandler)
	e.DELETE("/groups/:id/users", a.removeUsersFromGroup)
	e.PUT("/policies/:id/permissions", a.addPermissionsToPolicy)
	e.GET("/policies/:id/permissions", a.getPermissionsByPolicyHandler)
	e.DELETE("/policies/:id/permissions", a.removePermissionsFromPolicy)
	e.PUT("/users/:id/policies", a.attachPoliciesByUser)
	e.GET("/users/:id/policies", a.getPoliciesByUserHandler)
	e.DELETE("users/:id/policies", a.detachPoliciesByUser)
	e.PUT("/groups/:id/policies", a.attachPoliciesByGroup)
	e.GET("/groups/:id/policies", a.getPoliciesByGroupHandler)
	e.DELETE("/groups/:id/policies", a.detachPoliciesByGroup)

	e.PATCH("/permissions/:id", a.updatePermission)
	//e.DELETE("/permissions/:id", a.removePermission)

	e.POST("/check_permissions", a.userPermissions)
	e.Logger.Fatal(e.Start(":1535"))
}
