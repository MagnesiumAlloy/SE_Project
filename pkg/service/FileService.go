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

func checkIsDir(path, name string, isRoot bool, isBin bool) error {
	if isRoot {
		if err := handler.NewFileHandler(&model.Data{Path: path, Name: name, InBin: isBin, Type: model.Dir}).CheckTargetExist(); err != nil {
			return err
		}
		return nil
	} else {
		return handler.SysCheckIsDir(path + name + "/")
	}

}

func ReadDir(path, name string, isRoot bool, isBin bool) ([]model.Data, error) {
	var res []model.Data
	if err := checkIsDir(path, name, isRoot, isBin); err != nil {
		return nil, err
	}
	if isRoot {
		if isBin {
			if res, err = handler.NewFileHandler(&model.Data{BinPath: path + name + "/"}).ReadDir(isBin); err != nil {
				return nil, err
			}
		} else {
			if res, err = handler.NewFileHandler(&model.Data{Path: path + name + "/"}).ReadDir(isBin); err != nil {
				return nil, err
			}
		}
		return res, nil
	} else {
		return handler.SysReadDir(path + name + "/")
	}
}

func ReadAllFileAndDir(root string) ([]model.Data, error) {
	res := []model.Data{}
	err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
		if info.IsDir() {
			name := ""
			path := "/"
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
			obj, _ := handler.SysReadFileInfo(name, root+path)
			obj.Path = path
			res = append(res, *obj)

		} else {
			name := ""
			path := "/"
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
			obj, _ := handler.SysReadFileInfo(name, root+path)
			obj.Path = path
			res = append(res, *obj)
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
		if isBin {
			if err := handler.NewFileHandler(&model.Data{Name: name, BinPath: path, InBin: isBin}).CheckTargetExist(); err != nil {
				return err
			}
		} else {
			if err := handler.NewFileHandler(&model.Data{Name: name, Path: path, InBin: isBin}).CheckTargetExist(); err != nil {
				return err
			}
		}

	} else {
		return handler.SysCheckFileExist(path + name)
	}
	return nil
}
func Recover(srcName, srcPath, desPath string) error {
	if err = checkFileExist(srcName, srcPath, true); err != nil {
		return err
	}
	if obj, err = handler.NewFileHandler(&model.Data{Name: srcName, Path: srcPath, InBin: false}).GetTarget(); err != nil {
		return err
	}
	if err = handler.SysCopy(model.Root+obj.Path+obj.Name, desPath+obj.Name, obj.ModTime, obj.CreatTime); err != nil {
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
	if obj, err = handler.NewFileHandler(&model.Data{Name: desName, Path: desName, InBin: false}).GetTarget(); err != nil {
		return err
	}
	if err = handler.SysCompare(srcPath+srcPath, model.Root+obj.Path+obj.Name); err != nil {
		return err
	}
	return nil
}

func Delete(name, path string) error {
	if err := checkFileExist(name, path, true); err != nil {
		return err
	}
	if util.GetTargetType(name) == model.Dir {
		var files []model.Data
		if files, err = handler.NewFileHandler(&model.Data{Path: path + name + "/"}).GetAllInDir(); err != nil {
			return err
		}
		for i := range files {
			if files[i].BinPath, err = util.PrefixReplace(path, "/", files[i].BinPath); err != nil {
				return err
			}
			files[i].InBin = true
			if err = handler.NewFileHandler(&files[i]).MoveToBin(); err != nil {
				return err
			}
		}

	}
	obj, err := handler.NewFileHandler(&model.Data{Name: name, Path: path}).GetTarget()
	if err != nil {
		return err
	}
	obj.BinPath = "/"
	obj.InBin = true
	if err = handler.NewFileHandler(obj).MoveToBin(); err != nil {
		return err
	}
	if err := handler.SysMove(model.Root+path+name, model.Bin+path+name); err != nil {
		return err
	}
	return nil
}

func Backup(srcName, srcPath, desPath string) error {
	if err := checkFileExist(srcName, srcPath, false); err != nil {
		return err
	}
	if obj, err = handler.SysReadFileInfo(srcName, srcPath); err != nil {
		return err
	}
	obj.Path = desPath
	if err := handler.NewFileHandler(obj).Backup(); err != nil {
		return err
	}
	if err := handler.SysCopy(srcPath+srcName, model.Root+desPath+srcName, time.Now(), time.Now()); err != nil {
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
	if err := handler.NewFileHandler(&model.Data{Name: name, Path: path, InBin: true}).Clean(); err != nil {
		return err
	}
	if err := handler.SysCleanFile(model.Bin + path + name); err != nil {
		return err
	}
	return nil
}

func Recycle(name, path string) error {
	if err := checkFileExist(name, path, true, true); err != nil {
		return err
	}

	if util.GetTargetType(name) == model.Dir {
		if err := handler.NewFileHandler(&model.Data{Path: path + name + "/"}).Recycle(); err != nil {
			return err
		}
	}

	if err := handler.NewFileHandler(&model.Data{Name: name, Path: path}).Recycle(); err != nil {
		return err
	}

	if err := handler.SysMove(model.Bin+path+name, model.Root+path+name); err != nil {
		return err
	}

	return nil
}
