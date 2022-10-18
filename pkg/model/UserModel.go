package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName   string
	Password   string
	Salt       string
	UserType   string
	TotalSpace uint64
	UsedSpace  uint64
}
