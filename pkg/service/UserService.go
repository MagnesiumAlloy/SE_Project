package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Login(userName string, password string, userType string) error {
	var salted_pwd string
	var err error
	if err = handler.NewUserHandler(&model.User{UserName: userName, UserType: userType}).CheckUserExist(); err != nil {
		return errors.New("wrong user name or password")
	}

	if salted_pwd, err = handler.NewUserHandler(&model.User{UserName: userName, UserType: userType}).GetSaltedPassword(); err != nil {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(salted_pwd), []byte(password)); err != nil {
		return errors.New("wrong user name or password")
	}

	return nil
}
