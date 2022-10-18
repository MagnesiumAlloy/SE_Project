package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName   string `form:"username"`
	Password   string `form:"password"`
	Salt       string
	UserType   string
	TotalSpace uint64
	UsedSpace  uint64
}
