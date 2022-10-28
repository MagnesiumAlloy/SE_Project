package handler

import (
	"SE_Project/pkg/model"
	"SE_Project/pkg/util"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"syscall"
	"time"
)

var err error

func SysCheckIsDir(path string) error {
	folderinfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	//log.Println(folderinfo)
	//check is a dir
	if !folderinfo.IsDir() {
		return errors.New("target is not a dir")
	}
	return nil
}

func SysCheckFileExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	return nil
}

func SysReadDir(path string) ([]model.Data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	files, err := f.ReadDir(0)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var result []model.Data
	for _, file := range files {
		log.Println(file)
		info, err := file.Info()
		if err != nil {
			return nil, err
		}
		Type := model.Dir
		if !file.IsDir() {
			Type = util.GetTargetType(file.Name())
		}
		result = append(result, model.Data{
			Name: file.Name(),
			Type: Type,
			Size: uint64(info.Size()),
			Path: path})

	}
	return result, nil
}

func SysReadFileInfo(name, path string) (*model.Data, error) {
	res := model.Data{Name: name, Path: path}
	if err := SysCheckIsDir(path + name); err != nil {
		res.Type = util.GetTargetType(name)
	} else {
		res.Type = model.Dir
	}
	info, err := os.Stat(path + name)
	if err != nil {
		return nil, err
	}
	res.Size = uint64(info.Size())
	res.ModTime = info.ModTime()
	res.CreatTime = time.Unix(info.Sys().(*syscall.Stat_t).Ctim.Sec, 0)
	return &res, nil
}

func SysCopy(srcPath, desPath string, modTime time.Time, CreatTime time.Time) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	desFile, err := os.OpenFile(desPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = io.Copy(desFile, srcFile)
	defer desFile.Close()
	defer srcFile.Close()
	return err
}

func SysGetFileMd5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	md5h := md5.New()
	io.Copy(md5h, file)

	return hex.EncodeToString(md5h.Sum(nil)), nil
}

func SysCompare(srcPath, desPath string) error {
	var srcMd5, desMd5 string
	if srcMd5, err = SysGetFileMd5(srcPath); err != nil {
		return err
	}
	if desMd5, err = SysGetFileMd5(srcPath); err != nil {
		return err
	}
	if srcMd5 != desMd5 {
		return errors.New("not the same")
	}
	return nil
}

func SysCleanFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func SysMove(srcPath, desPath string) error {
	if err := os.Rename(srcPath, desPath); err != nil {
		return err
	}
	return nil
}
