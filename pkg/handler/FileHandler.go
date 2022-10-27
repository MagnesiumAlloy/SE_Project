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

func (fileHandler *FileHandler) CheckTargetExist() error {
	var result model.Data
	if err := fileHandler.Where(fileHandler.Obj).First(&result); err != nil {
		return err.Error
	}
	return nil
}

func (fileHandler *FileHandler) GetTarget() (*model.Data, error) {
	var result model.Data
	if err := fileHandler.Where(fileHandler.Obj).First(&result); err != nil {
		return nil, err.Error
	}
	return &result, nil
}

func (fileHandler *FileHandler) Delete() error {
	if err := fileHandler.Model(fileHandler.Obj).Update("isdeleted", "true").Error; err != nil {
		return err
	}
	return nil
}
func (fileHandler *FileHandler) Backup() error {
	var res, qry model.Data
	qry.Name = fileHandler.Obj.Name
	qry.Path = fileHandler.Obj.Path
	qry.IsDeleted = fileHandler.Obj.IsDeleted
	if fileHandler.Where(&qry).First(&res); err != nil {
		if err := fileHandler.Create(&fileHandler.Obj).Error; err != nil {
			return err
		}
	} else {
		if err := fileHandler.Model(&res).Updates(&fileHandler.Obj).Error; err != nil {
			return err
		}
	}
	return nil
}

func (fileHandler *FileHandler) Clean() error {

	return nil
}

func (fileHandler *FileHandler) Recycle() error {
	return nil
}
