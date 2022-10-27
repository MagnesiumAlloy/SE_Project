package model

import (
	"time"

	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	Path      string    `form:"path"`
	Size      uint64    `form:"size"`
	ZipSize   uint64    `form:"zipsize"`
	Type      string    `form:"type"`
	Name      string    `form:"name"`
	UserId    uint64    `form:"userid"`
	IsDeleted bool      `form:"isdeleted"`
	ModTime   time.Time `form:"modtime"`
	CreatTime time.Time `form:"creattime"`
	encrypted bool      `form:"encrypt"`
	key       string    `form"key"`
}
