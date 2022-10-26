package main

import (
	"SE_Project/router"
	//"golang.org/x/crypto/bcrypt"
)

func main() {

	init.initDB()
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
