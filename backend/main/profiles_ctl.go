package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
)

func NewID() (id string) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for tries := CFG.Db.IdTries; tries > 0; tries-- {
		p, err := password.Generate(CFG.Db.IdLen, rnd.Intn(CFG.Db.IdLen-3)+1, 0, false, true)
		eh(err)

		id = strings.ToUpper(p)
		DBC.EnsureConn()
		var c int
		DBC.C.Table("profiles").Where("id = ?", id).Count(&c)
		if c < 1 {
			return
		}
	}
	eh(errors.New("can't generate unique token"))
	return
}

//noinspection GoUnusedExportedFunction
func IdMiddleware(c *gin.Context) {
	id := c.Param("id")
	m, err := regexp.MatchString(fmt.Sprintf("[A-Z0-9]{%d}", CFG.Db.IdLen), id)
	eh(err)
	if !m {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Set("id", strings.ToUpper(id))
	c.Next()
}

//noinspection GoUnusedExportedFunction
func PhoneMiddleware(c *gin.Context) {
	phone := c.Param("phone")
	m, err := regexp.MatchString(`\d{10}`, phone)
	eh(err)
	if !m {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Set("phone", phone)
	c.Next()
}
