package model

import "gorm.io/gorm"

type ObjectPointer struct {
	gorm.Model
	Path      string `form:"path"`
	Size      uint64 `form:"size"`
	ZipSize   uint64 `form:"zipsize"`
	Type      string `form:"type"`
	Name      string `form:"name"`
	UserId    uint64 `form:"userid"`
	IsDeleted bool   `form:"isdeleted"`
}

type Folder struct {
	ObjectPointer
}

type File struct {
	ObjectPointer
}
