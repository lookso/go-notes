package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var _db *gorm.DB

func init() {
	_db, err := gorm.Open("mysql", "root:lookfor@Dny8$@/activity?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	_db.DB().SetConnMaxLifetime(10)
	_db.DB().SetMaxOpenConns(100)
	_db.DB().SetMaxIdleConns(20)
	_db.LogMode(true)
	fmt.Printf("mysql Gorm Connect Success\n")
}
func GetDB() *gorm.DB {
	return _db
}
