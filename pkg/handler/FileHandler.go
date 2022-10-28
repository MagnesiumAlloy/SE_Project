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
func (fileHandler *FileHandler) ReadDir(isBin bool) ([]model.Data, error) {
	var result []model.Data
	if err := fileHandler.Debug().Where(fileHandler.Obj).Where("in_bin", isBin).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (fileHandler *FileHandler) CheckTargetExist() error {
	var result model.Data
	if err := fileHandler.Where(fileHandler.Obj).Where("in_bin", fileHandler.Obj.InBin).First(&result).Error; err != nil {
		return err
	}
	return nil
}

func (fileHandler *FileHandler) GetTarget() (*model.Data, error) {
	var result model.Data
	if err := fileHandler.Where(fileHandler.Obj).Where("in_bin", fileHandler.Obj.InBin).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (fileHandler *FileHandler) GetAllInDir() ([]model.Data, error) {
	var result []model.Data
	if err := fileHandler.Debug().Where("path LIKE ?", fileHandler.Obj.Path+"%").Where("in_bin", fileHandler.Obj.InBin).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (fileHandler *FileHandler) Update() error {
	if err := fileHandler.Debug().Save(fileHandler.Obj).Error; err != nil {
		return err
	}
	return nil
}

func (fileHandler *FileHandler) MoveToBin() error {
	var qry model.Data
	var qryRes []model.Data
	qry.Path = fileHandler.Obj.Path
	qry.Name = fileHandler.Obj.Name
	qry.InBin = true
	if err := fileHandler.Where(&qry).Find(&qryRes).Error; err == nil {
		for _, obj := range qryRes {
			if err := fileHandler.Delete(&obj).Error; err != nil {
				return err
			}
		}
	}
	return fileHandler.Update()
}

func (fileHandler *FileHandler) Backup() error {
	var res, qry model.Data
	qry.Name = fileHandler.Obj.Name
	qry.Path = fileHandler.Obj.Path
	if err := fileHandler.Where(&qry).Where("in_bin", false).First(&res).Error; err != nil {
		if err := fileHandler.Model(&res).Updates(fileHandler.Obj).Error; err != nil {
			return err
		}
	} else {
		if err := fileHandler.Create(&fileHandler.Obj).Error; err != nil {
			return err
		}
	}
	return nil
}

func (fileHandler *FileHandler) Clean() error {
	if err := fileHandler.Where(fileHandler.Obj).Delete(&model.Data{}).Error; err != nil {
		return err
	}
	return nil
}

func (fileHandler *FileHandler) Recycle() error {
	if err := fileHandler.Debug().Table("data").Where(fileHandler.Obj).Update("in_bin", false).Error; err != nil {
		return err
	}
	return nil
}
