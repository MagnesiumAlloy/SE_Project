package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Login(userName string, password string, userType string) (uint, error) {
	var salted_pwd string
	var userid uint
	var err error
	if userid, err = handler.NewUserHandler(&model.User{UserName: userName, UserType: userType}).CheckUserExist(); err != nil {
		return 0, errors.New("wrong user name or password")
	}

	if salted_pwd, err = handler.NewUserHandler(&model.User{UserName: userName, UserType: userType}).GetSaltedPassword(); err != nil {
		return 0, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(salted_pwd), []byte(password)); err != nil {
		return 0, errors.New("wrong user name or password")
	}

	return userid, nil
}

func AdminLogin(pwd string) error {

	if pwd == "admin" {
		return nil
	}
	return errors.New("wrong password")
}
func Register(userName string, password string) error {
	var salted_pwd []byte
	var err error
	//check if exist
	if _, err = handler.NewUserHandler(&model.User{UserName: userName}).CheckUserExist(); err == nil {
		return errors.New("username already exists!")
	}
	if salted_pwd, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return err
	}
	if err = handler.NewUserHandler(&model.User{UserName: userName, UserType: model.NormalUserType, Password: string(salted_pwd)}).Register(); err != nil {
		return err
	}
	return nil
}
