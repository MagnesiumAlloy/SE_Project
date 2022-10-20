package Controller

import (
	svc "SE_Project/pkg/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ListFiles(c *gin.Context) {
	path := c.Query("path")
	//pre check
	//check path
	if err := svc.CheckIsDir(path); err != nil {
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

}
