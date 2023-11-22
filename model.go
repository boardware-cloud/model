package model

import (
	"fmt"

	"github.com/boardware-cloud/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = NewConnection(config.GetString("db.user"), config.GetString("db.password"), config.GetString("db.host"), config.GetString("db.port"), config.GetString("db.database"))
		if err != nil {
			panic(err)
		}
	}
	return db
}

func NewConnection(user, password, host, port, database string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
