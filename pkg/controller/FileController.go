package Controller

import (
	"SE_Project/pkg/model"
	svc "SE_Project/pkg/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var err error

func ListFiles(c *gin.Context) {
	path := c.Query("path")
	var result []model.Data
	//pre check
	//check path
	if err = svc.CheckIsDir(path); err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  "path doesn't exist.",
			})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "path is not a dir",
		})
		return
	}
	//read dir
	if result, err = svc.ReadDir(path); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "read error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   result,
	})
}
