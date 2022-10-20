package model

import "gorm.io/gorm"

type ObjectPointer struct {
	gorm.Model
	Path string
	Size uint64
	Type string
	Name string
}

type Folder struct {
	ObjectPointer
}

type File struct {
	ObjectPointer
}
