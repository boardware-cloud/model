package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewConnection(user, password, host, port, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)
	var err error
	if db == nil {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	return db, err
}
