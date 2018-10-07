package db

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/jinzhu/gorm"
	"log"
)

var SqlDB *gorm.DB
func InitDB() {
	var err error
	SqlDB, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/av?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	SqlDB.LogMode(true)
	SqlDB.DB().SetMaxOpenConns(1000)
	SqlDB.DB().SetMaxIdleConns(500)
	SqlDB.DB().SetConnMaxLifetime(60*time.Second)
	SqlDB.DB().Ping()
	//SqlDB.

}