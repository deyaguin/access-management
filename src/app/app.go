package app

import (
	"app/api"
	"app/db"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type App struct {

}

func (app *App) Init() {
	//pgDB, err := db.SqlDBCreator("postgres", "host=localhost user=accessControl dbname=accesscontrol password=agryz2010")
	sqliteDB, err := db.SqlDBCreator("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		//pgDB,
		sqliteDB,
	}
	api.Init()
}
