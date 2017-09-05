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

	usersService := services.NewUsersService(pgDB)
	groupsService := services.NewGroupsService(pgDB)
	policiesService := services.NewPoliciesService(pgDB)
	permissionsService := services.NewPermissionsService(pgDB)
	permissionsCheckService := services.NewPermissionsCheckService(pgDB)
	actionsService := services.NewActionsService(pgDB)

	api.NewAPI(
		usersService,
		groupsService,
		policiesService,
		permissionsService,
		permissionsCheckService,
		actionsService,
		":1535",
	)
}
