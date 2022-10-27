package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"SE_Project/pkg/util"
	"errors"
	"log"
	"os"
	"path/filepath"
)

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
		result = append(result, model.Data{
			Name: file.Name(),
			Type: util.GetTargetType(file.Name(), file.IsDir()),
			Size: uint64(info.Size()),
			Path: ""})

	}
	return result, nil
}

func CheckIsDir(path string, isRoot bool, isBin bool) error {
	if isRoot {
		if err := handler.NewFileHandler(&model.Data{Path: path, IsDeleted: isBin}).CheckIsDir(); err != nil {
			return err
		}
		return nil
	} else {
		return SysCheckIsDir(path)
	}

}

func ReadDir(path string, isRoot bool, isBin bool) ([]model.Data, error) {
	if isRoot {
		result, err := handler.NewFileHandler(&model.Data{Path: path, IsDeleted: isBin}).ReadDir()
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return SysReadDir(path)
	}
}

func ReadAllFileAndDir(root string) ([]model.Data, error) {
	var res []model.Data
	err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
		if info.IsDir() {
			name := ""
			path := ""
			for i, fd := len(file)-1, 0; i >= len(root); i-- {
				if fd == 0 && file[i] == '/' {
					fd = 1
				} else {
					if fd == 1 {
						path = string(file[i]) + path
					} else {
						name = string(file[i]) + name
					}
				}
			}
			res = append(res, model.Data{Path: path, Name: name, Type: model.Dir, Size: uint64(info.Size())})

		} else {
			name := ""
			path := ""
			for i, fd := len(file)-1, 0; i >= len(root); i-- {
				if fd == 0 && file[i] == '/' {
					fd = 1
				} else {
					if fd == 1 {
						path = string(file[i]) + path
					} else {
						name = string(file[i]) + name
					}
				}
			}
			res = append(res, model.Data{Path: path, Name: name, Type: util.GetTargetType(name, false), Size: uint64(info.Size())})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Recover(srcName, srcPath, desPath string) error {
	return nil
}

func Compare(srcName, srcPath, desName, desPath string) error {
	return nil
}

func Delete(Name, Path string) error {
	return nil
}

func Backup(srcName, srcPath, desPath string) error {
	return nil
}

func BackupWithKey(srcName, srcPath, desPath, key string) error {
	return nil
}

func Clean(name, path string) error {
	return nil
}

func Recycle(name, path string) error {
	return nil
}
