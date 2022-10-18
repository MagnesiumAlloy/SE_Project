package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"errors"

	"gorm.io/gorm"
)

func Login(loginForm *model.User) error {
	if err := handler.NewUserHandler(loginForm).Login(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("invalid user name or password")
		}
		return err
	}
	return nil
}
