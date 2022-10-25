package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	ID         string     `json:"id" gorm:"primary_key;type:varchar(10)"`
	NotifiedAt *time.Time `json:"notified_at"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Profile    JSON       `json:"profile" gorm:"type:json"`
}

func (s *Server) Add(c *gin.Context) {
	var r Profile
	DBC.EnsureConn()
	eh(c.BindJSON(&r.Profile))
	r.ID = NewID()
	if dbEhDub(DBC.C.Create(r).Error) {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	c.String(http.StatusCreated, "%s", r.ID)
}
