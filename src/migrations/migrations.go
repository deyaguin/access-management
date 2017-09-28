package migrations

import (
	"gitlab/nefco/access-management-system/src/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"gopkg.in/gormigrate.v1"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Migration struct {
}

var migrate1 = &gormigrate.Migration{
	ID: "0",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(
			&models.Service{},
			&models.Action{},
			&models.User{},
			&models.Group{},
			&models.Policy{},
			&models.Permission{},
		).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable(
			&models.Service{},
			&models.Action{},
			&models.User{},
			&models.Group{},
			&models.Policy{},
			&models.Permission{},
		).Error
	},
}

func (migration *Migration) Init() {
	pgDB, err := gorm.Open("postgres", "host=localhost user=access-management dbname=access-management password=12345")
	//sqliteDB, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	if err = pgDB.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	//if err = sqliteDB.DB().Ping(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//sqliteDB.LogMode(true)
	pgDB.LogMode(true)

	m := gormigrate.New(pgDB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrate1,
	})
	//m := gormigrate.New(sqliteDB, gormigrate.DefaultOptions, []*gormigrate.Migration{
	//	migrate1,
	//})

	if err = m.Migrate(); err != nil {
		log.Fatal(err)
	}

	log.Print("migration successfully")

}
