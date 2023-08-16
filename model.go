package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(user, password, host, port, database string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

type MayBe[T any] struct {
	Data *T
}

func (m MayBe[T]) Just(f func(data T)) MayBe[T] {
	if m.Data != nil {
		f(*m.Data)
	}
	return m
}

func (m MayBe[T]) Nothing(f func()) MayBe[T] {
	if m.Data == nil {
		f()
	}
	return m
}
