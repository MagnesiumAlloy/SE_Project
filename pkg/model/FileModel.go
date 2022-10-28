package model

import (
	"time"

	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	Name      string    `form:"name"`
	Path      string    `form:"path"`
	Type      string    `form:"type"`
	Size      uint64    `form:"size"`
	ZipSize   uint64    `form:"zipsize"`
	UserId    uint64    `form:"userid"`
	InBin     bool      `form:"inbin"`
	BinPath   string    `form:"binpath"`
	ModTime   time.Time `form:"modtime"`
	CreatTime time.Time `form:"creattime"`
	encrypted bool      `form:"encrypt"`
	key       string    `form"key"`
}
