package main

import (
	"SE_Project/internal/model"
	"SE_Project/internal/router"
	"os/user"
)

func main() {

	user, _ := user.Current()
	model.Bin = user.HomeDir + "/Cloud_Bin"
	model.Root = user.HomeDir + "/Cloud_Backup"
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
