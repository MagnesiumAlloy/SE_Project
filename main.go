package main

import (
	"SE_Project/router"
)

func main() {
	//handler.GetDb().AutoMigrate(&model.User{})
	//handler.GetDb().Create(&model.User{UserName: "admin", Password: "123"})

	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
