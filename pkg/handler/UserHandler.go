package handler

import (
	"SE_Project/pkg/model"

	"gorm.io/gorm"
)

type UserHandler struct {
	*gorm.DB
	User *model.User
}

func NewUserHandler(user *model.User) *UserHandler {
	return &UserHandler{GetDB(), user}
}

func (userHandler *UserHandler) Login() error {
	if err := userHandler.Where(userHandler.User).First(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (userHandler *UserHandler) CheckUserExist() (uint, error) {
	var res model.User
	if err := userHandler.Where(userHandler.User).First(&res).Error; err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (userHandler *UserHandler) GetSaltedPassword() (string, error) {
	var result model.User
	if err := userHandler.Select("Password").Where(userHandler.User).First(&result).Error; err != nil {
		return "", err
	}
	return result.Password, nil
}

func (userHandler *UserHandler) Register() error {
	if err := userHandler.Create(userHandler.User).Error; err != nil {
		return err
	}
	return userHandler.NewUserInit()
}

func (userHandler *UserHandler) NewUserInit() error {
	var obj *model.Data
	if obj, err = NewFileHandler(&model.Data{}).GetTarget(); err != nil {
		return err
	}
	obj.UserId = userHandler.User.ID
	obj.ID = 0
	if err := NewFileHandler(&model.Data{}).Create(obj).Error; err != nil {
		return err
	}
	obj.Path = ""
	obj.BinPath = "/"
	obj.InBin = true
	obj.ID = 0
	if err := NewFileHandler(&model.Data{}).Create(obj).Error; err != nil {
		return err
	}
	return nil
}
