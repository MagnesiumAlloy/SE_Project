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
	r.GET("/adminMain", func(c *gin.Context) {
		c.HTML(http.StatusOK, "adminMain.html", gin.H{
			"title": "adminMain",
		})
	})
	r.GET("/ScanUser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ScanUser.html", gin.H{
			"title": "ScanUser",
		})
	})
	r.GET("/DeleteFile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "DeleteFile.html", gin.H{
			"title": "DeleteFile",
		})
	})
	r.GET("/DeleteUser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "DeleteUser.html", gin.H{
			"title": "DeleteUser",
		})
	})
	r.GET("/Modify", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Modify.html", gin.H{
			"title": "Modify",
		})
	})
	r.GET("/myfile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "myfile.html", gin.H{
			"title": "myfile",
		})
	})
	r.GET("/recycle", func(c *gin.Context) {
		c.HTML(http.StatusOK, "recycle.html", gin.H{
			"title": "recycle",
		})
	})
	r.GET("/secret", func(c *gin.Context) {
		c.HTML(http.StatusOK, "secret.html", gin.H{
			"title": "secret",
		})
	})
	r.GET("/share", func(c *gin.Context) {
		c.HTML(http.StatusOK, "share.html", gin.H{
			"title": "share",
		})
	})
	r.GET("/zhuye", func(c *gin.Context) {
		c.HTML(http.StatusOK, "zhuye.html", gin.H{
			"title": "zhuye",
		})
	})
	r.GET("/select", func(c *gin.Context) {
		c.HTML(http.StatusOK, "select.html", gin.H{
			"title": "select",
		})
	})
	r.GET("/ScanFile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ScanFile.html", gin.H{
			"title": "ScanFile",
		})
	})
	r.GET("/normal", func(c *gin.Context) {
		c.HTML(http.StatusOK, "normal.html", gin.H{
			"title": "normal",
		})
	})

	r.POST("/Login", Controller.Login)

	r.POST("/register", Controller.Register)

	r.POST("/recover", Controller.Recover)

	r.POST("/compare", Controller.Compare)

	r.DELETE("/delete", Controller.Delete)

	r.GET("/readDir", Controller.ReadDir)

	r.POST("/backup", Controller.Backup)

	r.POST("/backupWithKey", Controller.BackupWithKey)

	r.DELETE("/clean", Controller.Clean)

	r.POST("/recycle", Controller.Recycle)
	return r
}
