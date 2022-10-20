package Controller

import (
	"SE_Project/pkg/model"
	svc "SE_Project/pkg/service"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginForm model.User
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	//login
	log.Println(loginForm.UserName)
	log.Println(loginForm.Password)
	log.Println(loginForm.UserType)

	if err := svc.Login(loginForm.UserName, loginForm.Password, loginForm.UserType); err != nil {
		c.JSON(401, gin.H{
			"status": 401,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
