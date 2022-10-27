package main

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"SE_Project/pkg/service"
	"SE_Project/router"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	//initDB()
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}

func initDB() {
	handler.GetDB().AutoMigrate(&model.User{})
	handler.GetDB().AutoMigrate(&model.Data{})
	//handler.GetDB().Raw("ALTER TABLE data ADD UNIQUE KEY(path, name);alter table data modify name varchar(256);")
	pwd, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	handler.GetDB().Create(&model.User{UserName: "user", Password: string(pwd), UserType: model.NormalUserType})
	pwd, _ = bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	handler.GetDB().Create(&model.User{UserName: "admin", Password: string(pwd), UserType: model.AdminType})

	res, _ := service.ReadAllFileAndDir(model.Root)
	for _, x := range res {
		handler.GetDB().Create(&x)
	}
}
