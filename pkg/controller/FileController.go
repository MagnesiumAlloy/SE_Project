package Controller

import (
	"SE_Project/pkg/model"
	svc "SE_Project/pkg/service"
	"SE_Project/pkg/validator"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

func Recover(c *gin.Context) {
	srcPath := c.PostForm("srcPath")
	desPath := c.PostForm("desPath")
	if err = validator.CheckNameAndPath([]string{}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Recover(srcPath, desPath); err != nil {
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
	srcPath := c.PostForm("srcPath")
	desPath := c.PostForm("desPath")
	if err = validator.CheckNameAndPath([]string{}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Compare(srcPath, desPath); err != nil {
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
	path := c.Query("path")
	if err = validator.CheckNameAndPath([]string{}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Delete(path); err != nil {
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
	path := c.Query("path")
	isRoot := c.Query("isroot") == "1"
	ID, _ := strconv.Atoi(c.Query("ID"))
	inBin := c.Query("inBin") == "1"
	log.Println(path, isRoot)
	var file []model.Data
	if err = validator.CheckNameAndPath([]string{}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if file, err = svc.ReadDir(path, &ID, isRoot, inBin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   file,
		"ID":     ID,
	})
}

func Backup(c *gin.Context) {
	srcPath := c.PostForm("srcPath")
	desPath := c.PostForm("desPath")
	if err = validator.CheckNameAndPath([]string{}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Backup(srcPath, desPath); err != nil {
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
	srcPath := c.PostForm("srcPath")
	desPath := c.PostForm("desPath")
	key := c.PostForm("key")
	if err = validator.CheckNameAndPath([]string{}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.BackupWithKey(srcPath, desPath, key); err != nil {
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

func Clean(c *gin.Context) {
	path := c.Query("path")
	if err = validator.CheckNameAndPath([]string{}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Clean(path); err != nil {
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
	path := c.PostForm("path")
	if err = validator.CheckNameAndPath([]string{}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Recycle(path); err != nil {
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
