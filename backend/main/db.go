package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct{ C *gorm.DB }

func (db *DB) Connect() {
	var err error
	db.C, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			CFG.Db.User, CFG.Db.Password, CFG.Db.Host, CFG.Db.Port, CFG.Db.DbName))
	eh(err, "db connect")
	// db.C.LogMode(CFG.Service.RunMode == "debug")
	db.C.SetLogger(LOG)
}

func (db *DB) EnsureConn() {
	eh(db.C.DB().Ping(), "db conn ping")
}

func (db *DB) Disconnect() {
	eh(db.C.Close(), "db disconnect")
}
