package main

import (
	"SE_Project/router"
)

func main() {
	//handler.GetDb().AutoMigrate(&model.User{})
	//pwd, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	//handler.GetDb().Create(&model.User{UserName: "user", Password: string(pwd), UserType: model.NormalUserType})
	//pwd, _ = bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	//handler.GetDb().Create(&model.User{UserName: "admin", Password: string(pwd), UserType: model.AdminType})

	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
