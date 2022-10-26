package init

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"

	"golang.org/x/crypto/bcrypt"
)

func initDB() {
	handler.GetDB().AutoMigrate(&model.User{})
	handler.GetDB().AutoMigrate(&model.Data{})
	//user 123 || admin admin
	pwd, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	handler.GetDB().Create(&model.User{UserName: "user", Password: string(pwd), UserType: model.AdminType})
	pwd, _ = bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	handler.GetDB().Create(&model.User{UserName: "admin", Password: string(pwd), UserType: model.AdminType})

}
