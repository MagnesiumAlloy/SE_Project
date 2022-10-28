package Controller

import (
	"SE_Project/pkg/model"
	svc "SE_Project/pkg/service"
	"SE_Project/pkg/validator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err error

func Recover(c *gin.Context) {
	srcName := c.PostForm("srcName")
	srcPath := c.PostForm("srcPath")
	desPath := c.PostForm("desPath")
	if err = validator.CheckNameAndPath([]string{srcName}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Recover(srcName, srcPath, desPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	messgae_map := map[string]interface{}{
		"status": 200,
		"msg":    "提交成功",
	}
	c.JSON(http.StatusOK, messgae_map)
}

func Compare(c *gin.Context) {
	srcName := c.PostForm("srcName")
	srcPath := c.PostForm("srcPath")
	desName := c.PostForm("desName")
	desPath := c.PostForm("desPath")
	if err = validator.CheckNameAndPath([]string{srcName, desName}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Compare(srcName, srcPath, desName, desPath); err != nil {
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

func Delete(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Delete(name, path); err != nil {
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

func ReadDir(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")
	isRoot := c.Query("isroot") == "1"
	log.Println(name, path, isRoot)
	var file []model.Data
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if file, err = svc.ReadDir(path, name, isRoot, false); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   file,
	})
}

func Backup(c *gin.Context) {
	srcName := c.PostForm("srcName")
	srcPath := c.PostForm("srcPath")
	desPath := c.PostForm("desPath")
	if err = validator.CheckNameAndPath([]string{srcName}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Backup(srcName, srcPath, desPath); err != nil {
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

func BackupWithKey(c *gin.Context) {
	srcName := c.PostForm("srcName")
	srcPath := c.PostForm("srcPath")
	desPath := c.PostForm("desPath")
	key := c.PostForm("key")
	if err = validator.CheckNameAndPath([]string{srcName}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.BackupWithKey(srcName, srcPath, desPath, key); err != nil {
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

func ReadBin(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")
	var file []model.Data
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if file, err = svc.ReadDir(path, name, true, true); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   file,
	})
}

func Clean(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Clean(name, path); err != nil {
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

func Recycle(c *gin.Context) {
	name := c.PostForm("name")
	path := c.PostForm("path")
	if err = validator.CheckNameAndPath([]string{name}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Recycle(name, path); err != nil {
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
