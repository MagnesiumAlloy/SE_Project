package model

import "gorm.io/gorm"

type ObjectPointer struct {
	gorm.Model
	Path   string `form:"filepath"`
	Size   uint64
	Type   string `form:"filetype"`
	Name   string `form:"filename"`
	Auther string
}

type Folder struct {
	ObjectPointer
}

type File struct {
	ObjectPointer
}
