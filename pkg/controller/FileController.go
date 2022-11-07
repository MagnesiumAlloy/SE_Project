package Controller

import (
	"SE_Project/pkg/model"
	svc "SE_Project/pkg/service"
	"SE_Project/pkg/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

func Recover(c *gin.Context) {
	srcPath := c.PostForm("srcPath")
	desPath := c.PostForm("desPath")
	UserId, _ := strconv.Atoi(c.PostForm("UserId"))
	if err = validator.CheckNameAndPath([]string{}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Recover(srcPath, desPath, uint(UserId)); err != nil {
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
	UserId, _ := strconv.Atoi(c.PostForm("UserId"))
	if err = validator.CheckNameAndPath([]string{}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Compare(srcPath, desPath, uint(UserId)); err != nil {
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
	UserId, _ := strconv.Atoi(c.Query("UserId"))
	if err = validator.CheckNameAndPath([]string{}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Delete(path, uint(UserId)); err != nil {
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
	UserId, _ := strconv.Atoi(c.Query("UserId"))
	var file []model.Data
	if err = validator.CheckNameAndPath([]string{}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if file, err = svc.ReadDir(path, &ID, isRoot, inBin, uint(UserId)); err != nil {
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
	UserId, _ := strconv.Atoi(c.PostForm("UserId"))
	pack := c.PostForm("pack") == "1"
	if err = validator.CheckNameAndPath([]string{}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Backup(srcPath, desPath, "", uint(UserId), pack, false); err != nil {
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
	pack := c.PostForm("pack") == "1"
	UserId, _ := strconv.Atoi(c.PostForm("UserId"))

	if err = validator.CheckNameAndPath([]string{}, []string{srcPath, desPath}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Backup(srcPath, desPath, key, uint(UserId), pack, true); err != nil {
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
	UserId, _ := strconv.Atoi(c.Query("UserId"))

	if err = validator.CheckNameAndPath([]string{}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Clean(path, uint(UserId)); err != nil {
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
	UserId, _ := strconv.Atoi(c.PostForm("UserId"))

	if err = validator.CheckNameAndPath([]string{}, []string{path}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	if err = svc.Recycle(path, uint(UserId)); err != nil {
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
