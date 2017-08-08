package dataBase

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	Vendor string
	Url string
}

func (dataBase *DB) Connect() *gorm.DB {
	db, err := gorm.Open(dataBase.Vendor, dataBase.Url)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}
