package app

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab/nefco/accessControl/app/api"
	"gitlab/nefco/accessControl/app/storage"
)

type App struct {
}

func (app *App) Init() {
	pgDB, err := storage.SqlDBCreator("postgres", "host=localhost user=accessControl dbname=accesscontrol password=agryz2010")
	//sqliteDB, err := db.SqlDBCreator("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		pgDB,
		//sqliteDB,
	}
	api.NewApi()
}
