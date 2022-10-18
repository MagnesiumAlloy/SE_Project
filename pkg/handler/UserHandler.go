package handler

import (
	"SE_Project/pkg/model"

	"gorm.io/gorm"
)

type UserHandler struct {
	conn *gorm.DB
	User *model.User
}

func NewUserHandler(user *model.User) *UserHandler {
	return &UserHandler{GetDb(), user}
}
func (userHandler *UserHandler) Login() error {
	if err := userHandler.conn.Where(userHandler.User).First(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
