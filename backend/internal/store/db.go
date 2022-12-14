package store

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sethvargo/go-password/password"
)

func NewID(dbc *gorm.DB, tries int, idLen int) (id string, err error) {
	rnd := rand.New(rand.NewSource(time.Now().Unix())) //#nosec
	for t := tries; t > 0; t-- {
		p, err := password.Generate(idLen, rnd.Intn(idLen-3)+1, 0, false, true)
		if err != nil {
			return "", err
		}

		id = strings.ToUpper(p)

		var c int
		dbc.Table("profiles").Where("id = ?", id).Count(&c)
		if c < 1 {
			return id, nil
		}
	}

	return "", errors.New("can't generate unique token")
}

func Connect(cfg Config) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName))
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
