package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName   string `form:"username"`
	Password   string `form:"password"`
	UserType   string `form:"usertype"`
	TotalSpace uint64
	UsedSpace  uint64
}
