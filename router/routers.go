package router

import (
	Controller "SE_Project/pkg/controller"
	"SE_Project/pkg/model"
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
	r.GET("/normal", func(c *gin.Context) {
		c.HTML(http.StatusOK, "normal.html", gin.H{
			"title": "normal",
		})
	})

	r.POST("/Login", Controller.Login)
	/*
		data := []model.ObjectPointer{
			{Name: "file1", Type: "t1", Path: "p1", Auther: "a1"},
			{Name: "file1", Type: "t1", Path: "p1", Auther: "a1"},
		}

		r.GET("/fileData", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		})
	*/
	r.GET("/fileData", Controller.ListFiles)
	data1 := []model.ObjectPointer{
		{Name: "file2", Type: "t2", Path: "p2", UserId: 2},
		{Name: "file3", Type: "t3", Path: "p3", UserId: 3},
	}
	r.GET("/innerData", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"data": data1,
		})
	})
	return r
}
