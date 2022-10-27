package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"SE_Project/pkg/util"
	"os"
	"path/filepath"
	"time"
)

var err error
var obj *model.Data

func checkIsDir(path string, isRoot bool, isBin bool) error {
	if isRoot {
		if err := handler.NewFileHandler(&model.Data{Path: path, IsDeleted: isBin, Type: model.Dir}).CheckTargetExist(); err != nil {
			return err
		}
		return nil
	} else {
		return handler.SysCheckIsDir(path)
	}

}

func ReadDir(path string, isRoot bool, isBin bool) ([]model.Data, error) {
	if err := checkIsDir(path, isRoot, isBin); err != nil {
		return nil, err
	}
	if isRoot {
		result, err := handler.NewFileHandler(&model.Data{Path: path, IsDeleted: isBin}).ReadDir()
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return handler.SysReadDir(path)
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

func checkFileExist(name, path string, params ...bool) error {
	isRoot := params[0]
	if isRoot {
		isBin := false
		if len(params) > 1 {
			isBin = params[1]
		}
		if err := handler.NewFileHandler(&model.Data{Name: name, Path: path, IsDeleted: isBin}).CheckTargetExist(); err != nil {
			return err
		}
	} else {
		return handler.SysCheckFileExist(path + "/" + name)
	}
	return nil
}
func Recover(srcName, srcPath, desPath string) error {
	if err = checkFileExist(srcName, srcPath, true); err != nil {
		return err
	}
	if obj, err = handler.NewFileHandler(&model.Data{Name: srcName, Path: srcName, IsDeleted: false}).GetTarget(); err != nil {
		return err
	}
	if err = handler.SysCopy(model.Root+obj.Path+"/"+obj.Name, desPath+"/"+obj.Name, obj.ModTime, obj.CreatTime); err != nil {
		return err
	}
	return nil
}

func Compare(srcName, srcPath, desName, desPath string) error {
	if err := checkFileExist(srcName, srcPath, false); err != nil {
		return err
	}
	if err := checkFileExist(desName, desPath, true); err != nil {
		return err
	}
	if obj, err = handler.NewFileHandler(&model.Data{Name: desName, Path: desName, IsDeleted: false}).GetTarget(); err != nil {
		return err
	}
	if err = handler.SysCompare(srcPath+"/"+srcPath, model.Root+obj.Path+"/"+obj.Name); err != nil {
		return err
	}
	return nil
}

func Delete(name, path string) error {
	if err := checkFileExist(name, path, true); err != nil {
		return err
	}
	if err := handler.NewFileHandler(&model.Data{Name: name, Path: path, Type: model.Dir}).CheckTargetExist(); err == nil {
		if err := handler.NewFileHandler(&model.Data{Path: path + "/" + name}).Delete(); err != nil {
			return err
		}
	}
	if err := handler.NewFileHandler(&model.Data{Name: name, Path: path}).Delete(); err != nil {
		return err
	}
	return nil
}

func Backup(srcName, srcPath, desPath string) error {
	if err := checkFileExist(srcName, srcName, false); err != nil {
		return err
	}
	if obj, err = handler.SysReadFileInfo(srcName, srcPath); err != nil {
		return err
	}

	if err := handler.NewFileHandler(obj).Backup(); err != nil {
		return err
	}
	if err := handler.SysCopy(srcPath+"/"+srcName, model.Root+desPath, time.Now(), time.Now()); err != nil {
		return err
	}

	return nil
}

func BackupWithKey(srcName, srcPath, desPath, key string) error {
	if err := checkFileExist(srcName, srcName, false); err != nil {
		return err
	}
	//todo
	return nil
}

func Clean(name, path string) error {
	if err := checkFileExist(name, path, true, true); err != nil {
		return err
	}
	if err := handler.NewFileHandler(&model.Data{Name: name, Path: path, IsDeleted: true}).Clean(); err != nil {
		return err
	}
	if err := handler.SysCleanFile(model.Bin + path + "/" + name); err != nil {
		return err
	}
	return nil
}

func Recycle(name, path string) error {
	if err := checkFileExist(name, path, true, true); err != nil {
		return err
	}
	if err := handler.NewFileHandler(&model.Data{Name: name, Path: path}).Recycle(); err != nil {
		return err
	}

	if err := handler.SysCopy(model.Bin+path+"/"+name, model.Root+path+"/"+name, time.Now(), time.Now()); err != nil {
		return err
	}
	return nil
}
