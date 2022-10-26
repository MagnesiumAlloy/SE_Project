package handler

import (
	"SE_Project/pkg/model"

	"gorm.io/gorm"
)

type FileHandler struct {
	*gorm.DB
	Obj *model.Data
}

func NewFileHandler(Obj *model.Data) *FileHandler {
	return &FileHandler{GetDB(), Obj}
}
func (fileHandler *FileHandler) ReadDir() ([]model.Data, error) {
	var result []model.Data
	if err := fileHandler.Where(fileHandler.Obj).Find(&result); err != nil {
		return nil, err.Error
	}
	return result, nil
}
func (fileHandler *FileHandler) CheckIsDir() error {
	var result model.Data
	if err := fileHandler.Where(fileHandler.Obj).First(&result); err != nil {
		return err.Error
	}
	return nil
}
