package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab/nefco/access-management-system/src/api"
	"gitlab/nefco/access-management-system/src/services"
	"gitlab/nefco/access-management-system/src/storage"
)

func main() {

	//new(migrations.Migration).Init()
	pgDB := storage.SqlDBCreator(
		"postgres",
		"host=localhost user=accessControl dbname=accesscontrol password=agryz2010",
	)
	//pgDB := storage.SqlDBCreator("sqlite3", "test.db")

	userService := services.NewUserService(pgDB)
	groupService := services.NewGroupService(pgDB)
	policyService := services.NewPolicyService(pgDB)
	permissionService := services.NewPermissionService(pgDB)
	permissionsCheckService := services.NewPermissionsCheckService(pgDB)
	actionsService := services.ActionsService(pgDB)

	api.NewAPI(
		userService,
		groupService,
		policyService,
		permissionService,
		permissionsCheckService,
		actionsService,
		":1535",
	)
}
