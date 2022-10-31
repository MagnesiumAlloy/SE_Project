package handler

import (
	"SE_Project/pkg/model"
	"path/filepath"

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
	if err := fileHandler.Debug().Where(fileHandler.Obj).Where("in_bin", fileHandler.Obj.InBin).Find(&result).Error; err != nil {
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
	if fileHandler.Obj.InBin {
		if err := fileHandler.Where("bin_path LIKE ?", fileHandler.Obj.BinPath+"%").Where("in_bin", fileHandler.Obj.InBin).Where("user_id", fileHandler.Obj.UserId).Find(&result).Error; err != nil {
			return nil, err
		}
	} else {
		if err := fileHandler.Where("path LIKE ?", fileHandler.Obj.Path+"%").Where("in_bin", fileHandler.Obj.InBin).Where("user_id", fileHandler.Obj.UserId).Find(&result).Error; err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (fileHandler *FileHandler) Update() error {
	if err := fileHandler.Save(fileHandler.Obj).Error; err != nil {
		return err
	}
	return nil
}

func (fileHandler *FileHandler) MoveToBin() error {
	var qryRes []model.Data
	if err := fileHandler.Where("path", fileHandler.Obj.Path).Where("in_bin", true).Where("user_id", fileHandler.Obj.UserId).Find(&qryRes).Error; err == nil {
		for _, obj := range qryRes {
			if err := fileHandler.Delete(&obj).Error; err != nil {
				return err
			}
			if err := SysClean(filepath.Join(model.Bin, fileHandler.Obj.Path)); err != nil {
				return err
			}
		}
	}

	return fileHandler.Update()
}

func (fileHandler *FileHandler) Backup() error {
	var res model.Data

	if err := fileHandler.Where("path", fileHandler.Obj.Path).Where("in_bin", false).Where("user_id", fileHandler.Obj.UserId).First(&res).Error; err != nil {
		//record not found
		if err := fileHandler.Create(&fileHandler.Obj).Error; err != nil {
			return err
		}
	} else {
		if err := fileHandler.Model(&res).Updates(fileHandler.Obj).Error; err != nil {
			return err
		}
	}
	return nil
}

func (fileHandler *FileHandler) Clean() error {
	/*
		var res *model.Data
		if res, err = fileHandler.GetTarget(); err != nil {
			return err
		}
		if err := fileHandler.Delete(res).Error; err != nil {
			return err
		}
	*/
	if err := fileHandler.Delete(fileHandler.Obj).Error; err != nil {
		return err
	}
	return nil
}

func (fileHandler *FileHandler) Recycle() error {
	var qryRes []model.Data
	if err := fileHandler.Where("path", fileHandler.Obj.Path).Where("in_bin", false).Where("user_id", fileHandler.Obj.UserId).Find(&qryRes).Error; err == nil && len(qryRes) > 0 {
		for _, obj := range qryRes {
			if err := fileHandler.Delete(&obj).Error; err != nil {
				return err
			}
			if err := SysClean(filepath.Join(model.Root, fileHandler.Obj.Path)); err != nil {
				return err
			}
		}
	}

	return fileHandler.Update()
}
