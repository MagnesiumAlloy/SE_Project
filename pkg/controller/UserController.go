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
	log.Println(c.Param("username"))
	svc.Login(&loginForm)

	c.String(200, "Success")
}
