package handler

import (
	"SE_Project/pkg/model"
	"fmt"
	"time"

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
	if err := userHandler.NewUserInit(userHandler.User.ID); err != nil {
		return err
	}
	if err := SysNewUserDir(userHandler.User.ID); err != nil {
		return err
	}
	return nil
}

func (userHandler *UserHandler) NewUserInit(userID uint) error {
	obj := &model.Data{
		Name:    fmt.Sprint(userID),
		Path:    "/" + fmt.Sprint(userID),
		Type:    model.Dir,
		UserId:  userID,
		ModTime: time.Now(),
	}
	if err := NewFileHandler(&model.Data{}).Create(obj).Error; err != nil {
		return err
	}
	obj.BinPath = obj.Path
	obj.InBin = true
	obj.ID = 0
	if err := NewFileHandler(&model.Data{}).Create(obj).Error; err != nil {
		return err
	}
	return nil
}

func (userHandler *UserHandler) UpdatePassword() error {
	if err := userHandler.Model(userHandler.User).Update("password", userHandler.User.Password).Error; err != nil {
		return err
	}
	return nil
}

func (userHandler *UserHandler) ReadUser() ([]model.User, error) {
	var res []model.User
	if err := userHandler.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
