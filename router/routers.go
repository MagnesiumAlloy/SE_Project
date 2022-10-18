package router

import (
	Controller "SE_Project/pkg/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/html/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"title": "about",
		})
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "register",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "login",
		})
	})
	r.GET("/adminLogin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "adminLogin.html", gin.H{
			"title": "adminLogin",
		})
	})
	r.POST("/Login", Controller.Login)
	return r
}
