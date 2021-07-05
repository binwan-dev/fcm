package utils

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db          *gorm.DB
	dburi       string
	idleMaxConn int
	maxOpenConn int
)

func InitDb(uri string, idlemaxConn, maxopenConn int) {
	var err error
	Db, err = gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Panicf("Mysql connection failed! %s", err)
		panic(err)
	}

	if idlemaxConn < 0 {
		panic("idle connection must be greater than 0")
	}
	if maxopenConn < 0 {
		panic("max open connection must be greater than 0")
	}

	dburi = uri

	print(dburi)
}
