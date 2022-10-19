package Controller

import (
	"SE_Project/pkg/model"
	svc "SE_Project/pkg/service"

	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginForm model.User
	if err := c.ShouldBind(&loginForm); err != nil {
		return
	}
	//login
	log.Println(loginForm.UserName)
	log.Println(loginForm.Password)
	log.Println(loginForm.UserType)

	if err := svc.Login(loginForm.UserName, loginForm.Password, loginForm.UserType); err != nil {
		c.String(401, err.Error())
		return
	}
	c.String(200, "Success")
}
