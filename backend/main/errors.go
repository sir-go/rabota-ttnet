package main

import (
	"github.com/go-sql-driver/mysql"
)

func eh(err error, msg ...string) {
	if err == nil {
		return
	}
	LOG.Panic("panic err: ", err, " msg ", msg)
}

//noinspection GoUnusedFunction
func ehSkip(err error, msg ...string) {
	if err == nil {
		return
	}
	if len(msg) > 0 {
		LOG.Println(err, msg)
	} else {
		LOG.Println(err)
	}
}

func dbEhDub(err error) bool {
	if err == nil {
		return false
	}
	if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1062 {
		ehSkip(err)
		return true
	}
	eh(err)
	return false
}
