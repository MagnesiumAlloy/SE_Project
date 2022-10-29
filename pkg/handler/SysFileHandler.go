package handler

import (
	"SE_Project/pkg/model"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)

var err error

func SysCheckIsDir(path string) error {
	folderinfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !folderinfo.IsDir() {
		return errors.New("target is not a dir")
	}
	return nil
}

func SysCheckFileExist(path string) error {
	_, err := os.Stat(path)
	return err
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
	result := []model.Data{}
	for _, file := range files {
		//log.Println(file)
		info, err := file.Info()
		if err != nil {
			return nil, err
		}
		Type := model.Dir
		if !info.IsDir() {
			Type = filepath.Ext(file.Name())
		}
		result = append(result, model.Data{
			Name:      file.Name(),
			Type:      Type,
			Size:      uint64(info.Size()),
			Path:      filepath.Join(path, file.Name()),
			Perm:      uint32(info.Mode()),
			ZipSize:   uint64(info.Size()),
			Encrypted: false,
			ModTime:   info.ModTime(),
		})
	}
	return result, nil
}

//backup use this
func SysReadFileInfo(path string) (*model.Data, error) {
	Type := model.Dir
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		Type = filepath.Ext(path)
	}

	return &model.Data{
		Name:    info.Name(),
		Path:    path,
		Size:    uint64(info.Size()),
		ZipSize: uint64(info.Size()),
		ModTime: info.ModTime(),
		Perm:    uint32(info.Mode()),
		Type:    Type,
	}, nil
}

func SysModifyInfo(path string, perm uint32, modTime time.Time) error {
	if err := SysCheckFileExist(path); err != nil {
		return err
	}
	if err := os.Chtimes(path, time.Now(), modTime); err != nil {
		return err
	}
	if err := os.Chmod(path, os.FileMode(perm)); err != nil {
		return err
	}
	return nil
}

func sysCopyFile(srcPath, desPath string, perm uint32, modTime time.Time) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	desFile, err := os.OpenFile(desPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	if _, err = io.Copy(desFile, srcFile); err != nil {
		return err
	}
	if err = SysModifyInfo(desPath, perm, modTime); err != nil {
		return err
	}
	defer desFile.Close()
	defer srcFile.Close()
	return nil
}

func SysCopy(srcPath, desPath string, perm uint32, modTime time.Time) error {
	src, err := os.Stat(srcPath)
	if err != nil {
		return err
	}
	if src.IsDir() {
		path := filepath.Dir(filepath.Join(desPath, filepath.Base(srcPath)))
		if _, err := os.Stat(path); err != nil {
			if err := os.Mkdir(path, os.ModePerm); err != nil {
				return err
			}
		}
		if list, err := os.ReadDir(srcPath); err == nil {
			for _, item := range list {
				if err := SysCopy(filepath.Join(srcPath, item.Name()), filepath.Join(desPath, item.Name()), perm, modTime); err != nil {
					return err
				}
			}
		} else {
			return err
		}
	} else {
		path := filepath.Dir(desPath)
		if _, err := os.Stat(path); err != nil {
			if err := os.Mkdir(path, os.ModePerm); err != nil {
				return err
			}
		}
		return sysCopyFile(srcPath, desPath, perm, modTime)
	}
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
	if desMd5, err = SysGetFileMd5(desPath); err != nil {
		return err
	}
	if srcMd5 != desMd5 {
		return errors.New("not the same")
	}
	return nil
}

func sysCleanFile(path string) error {
	return os.Remove(path)
}

func SysClean(path string) error {
	src, err := os.Stat(path)
	if err != nil {
		return err
	}
	if src.IsDir() {
		if list, err := os.ReadDir(path); err == nil {
			for _, item := range list {
				if err := SysClean(filepath.Join(path, item.Name())); err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}
	return sysCleanFile(path)
}

func SysMove(srcPath, desPath string) error {
	/*
		if err := SysCheckFileExist(desPath); err == nil {
			if err := SysClean(desPath); err != nil {
				return err
			}
		}
	*/
	path := filepath.Dir(desPath)
	if _, err := os.Stat(path); err != nil {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
	}
	return os.Rename(srcPath, desPath)
}

func ReadAllFileAndDir(root string) ([]model.Data, error) {
	res := []model.Data{}
	err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
		if info.IsDir() {
			obj, _ := SysReadFileInfo(file)
			obj.Path = file
			res = append(res, *obj)
		} else {
			obj, _ := SysReadFileInfo(file)
			obj.Path = file
			res = append(res, *obj)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
