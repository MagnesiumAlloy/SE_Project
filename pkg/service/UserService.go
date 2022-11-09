package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Login(userName string, password string, userType string) (uint, error) {
	var saltedPassword string
	var userid uint
	var err error
	if userid, err = handler.NewUserHandler(&model.User{UserName: userName, UserType: userType}).CheckUserExist(); err != nil {
		return 0, errors.New("wrong user name or password")
	}

	if saltedPassword, err = handler.NewUserHandler(&model.User{UserName: userName, UserType: userType}).GetSaltedPassword(); err != nil {
		return 0, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(saltedPassword), []byte(password)); err != nil {
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
	var saltedPassword []byte
	var err error
	//check if exist
	if _, err = handler.NewUserHandler(&model.User{UserName: userName}).CheckUserExist(); err == nil {
		return errors.New("username already exists")
	}
	if saltedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return err
	}
	if err = handler.NewUserHandler(&model.User{UserName: userName, UserType: model.NormalUserType, Password: string(saltedPassword)}).Register(); err != nil {
		return err
	}
	return nil
}

func UpdatePassword(userID uint, password, newPassword string) error {
	var saltedPassword string
	user := &model.User{}
	user.ID = userID
	if saltedPassword, err = handler.NewUserHandler(user).GetSaltedPassword(); err != nil {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(saltedPassword), []byte(password)); err != nil {
		return errors.New("wrong password")
	}
	var saltedNewPassword []byte
	if saltedNewPassword, err = bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost); err != nil {
		return err
	}
	user.Password = string(saltedNewPassword)
	if err = handler.NewUserHandler(user).UpdatePassword(); err != nil {
		return err
	}
	return nil
}

func ReadUser() ([]model.User, error) {
	var users []model.User
	if users, err = handler.NewUserHandler(&model.User{}).ReadUser(); err != nil {
		return nil, err
	}
	return users, nil
}
