package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab/nefco/accessControl/app/api"
	"gitlab/nefco/accessControl/app/services"
	"gitlab/nefco/accessControl/app/storage"
)

func main() {
	pgDB, err := storage.SqlDBCreator("postgres", "host=localhost user=accessControl dbname=accesscontrol password=agryz2010")
	//sqliteDB, err := db.SqlDBCreator("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	userService := services.NewUserService(pgDB)
	groupService := services.NewGroupService(pgDB)
	policyService := services.NewPolicyService(pgDB)
	permissionService := services.NewPermissionService(pgDB)

	api.NewAPI(
		userService,
		groupService,
		policyService,
		permissionService,
	)

}
