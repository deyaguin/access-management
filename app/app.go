package app

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab/nefco/accessControl/app/api"
	"gitlab/nefco/accessControl/app/db"
)

type App struct {
	api api.Api
}

func (app *App) AppCreator() {
	pgDB, err := db.SqlDBCreator("postgres", "host=localhost user=accessControl dbname=accesscontrol password=agryz2010")
	//sqliteDB, err := db.SqlDBCreator("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		pgDB,
		//sqliteDB,
	}
	app.api = api
}

func (app *App) Init() {
	app.api.Init()
}
