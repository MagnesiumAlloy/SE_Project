package Controller

import (
	"SE_Project/pkg/model"
	svc "SE_Project/pkg/service"
	"SE_Project/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err error

func Recover(c *gin.Context) {
	srcName := c.Query("srcName")
	srcPath := c.Query("srcPath")
	desPath := c.Query("desPath")
	if err = validator.CheckNameAndPath([]string{srcName}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if err = svc.Recover(srcName, srcPath, desPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func Compare(c *gin.Context) {
	srcName := c.Query("srcName")
	srcPath := c.Query("srcPath")
	desName := c.Query("desName")
	desPath := c.Query("desPath")
	if err = validator.CheckNameAndPath([]string{srcName, desName}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if err = svc.Compare(srcName, srcPath, desName, desPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func Delete(c *gin.Context) {
	name := c.Query("Name")
	path := c.Query("Path")
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if err = svc.Delete(name, path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func ReadDir(c *gin.Context) {
	name := c.Query("Name")
	path := c.Query("Path")
	isRoot := c.Query("isroot") == "1"
	var file []model.Data
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if file, err = svc.ReadDir(name+"/"+path, isRoot, false); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   file,
	})
}

func Backup(c *gin.Context) {
	srcName := c.Query("srcName")
	srcPath := c.Query("srcPath")
	desPath := c.Query("desPath")
	if err = validator.CheckNameAndPath([]string{srcName}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if err = svc.Backup(srcName, srcPath, desPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func BackupWithKey(c *gin.Context) {
	srcName := c.Query("srcName")
	srcPath := c.Query("srcPath")
	desPath := c.Query("desPath")
	key := c.Query("key")
	if err = validator.CheckNameAndPath([]string{srcName}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if err = svc.BackupWithKey(srcName, srcPath, desPath, key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func ReadBin(c *gin.Context) {
	name := c.Query("Name")
	path := c.Query("Path")
	var file []model.Data
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if file, err = svc.ReadDir(name+"/"+path, true, true); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   file,
	})
}

func Clean(c *gin.Context) {
	name := c.Query("Name")
	path := c.Query("Path")
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if err = svc.Clean(name, path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func Recycle(c *gin.Context) {
	name := c.Query("Name")
	path := c.Query("Path")
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	if err = svc.Recycle(name, path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
