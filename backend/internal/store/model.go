package store

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

func AddProfile(c *gin.Context) {
	var r Profile

	dbConfig, _ := c.Get("dbConfig")
	cfg := dbConfig.(Config)
	dbc, err := Connect(cfg)
	if err != nil {
		panic(err)
	}

	if err = c.ShouldBindJSON(&r.Profile); err != nil {
		panic(err)
	}
	if r.ID, err = NewID(dbc, cfg.IdTries, cfg.IdLen); err != nil {
		panic(err)
	}

	if err = dbc.Create(r).Error; err != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	c.String(http.StatusCreated, "%s", r.ID)
}
