package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"os"
	"path/filepath"
	"time"
)

var err error
var obj *model.Data

func checkIsDir(path string, isRoot bool, inBin bool) error {
	if isRoot {
		if err := handler.NewFileHandler(&model.Data{Path: path, InBin: inBin, Type: model.Dir}).CheckTargetExist(); err != nil {
			return err
		}
		return nil
	} else {
		return handler.SysCheckIsDir(path)
	}

}

func ReadDir(path string, ID *int, isRoot bool, inBin bool) ([]model.Data, error) {
	var res []model.Data
	if err := checkIsDir(path, isRoot, inBin); err != nil {
		return nil, err
	}
	if isRoot {
		if inBin {
			if path == "/" {
				if obj, err := handler.NewFileHandler(&model.Data{BinPath: path, InBin: inBin}).GetTarget(); err != nil {
					return nil, err
				} else {
					*ID = int(obj.ID)
				}
			}
			if res, err = handler.NewFileHandler(&model.Data{PID: uint(*ID), InBin: inBin}).ReadDir(); err != nil {
				return nil, err
			}
		} else {
			if path == "/" {
				if obj, err := handler.NewFileHandler(&model.Data{Path: path, InBin: inBin}).GetTarget(); err != nil {
					return nil, err
				} else {
					*ID = int(obj.ID)
				}
			}
			if res, err = handler.NewFileHandler(&model.Data{PID: uint(*ID), InBin: inBin}).ReadDir(); err != nil {
				return nil, err
			}
		}
		return res, nil
	} else {
		return handler.SysReadDir(path)
	}
}

func checkFileExist(path string, isRoot bool, inBin bool) error {
	if isRoot {
		if inBin {
			if err := handler.NewFileHandler(&model.Data{BinPath: path, InBin: inBin}).CheckTargetExist(); err != nil {
				return err
			}
		} else {
			if err := handler.NewFileHandler(&model.Data{Path: path, InBin: inBin}).CheckTargetExist(); err != nil {
				return err
			}
		}

	} else {
		return handler.SysCheckFileExist(path)
	}
	return nil
}
func Recover(srcPath, desPath string) error {
	if err = checkFileExist(srcPath, true, false); err != nil {
		return err
	}
	if obj, err = handler.NewFileHandler(&model.Data{Path: srcPath, InBin: false}).GetTarget(); err != nil {
		return err
	}
	if err = handler.SysCopy(filepath.Join(model.Root, obj.Path), filepath.Join(desPath, filepath.Base(srcPath)), obj.Perm, obj.ModTime); err != nil {
		return err
	}
	return nil
}

func Compare(srcPath, desPath string) error {
	if err := checkFileExist(srcPath, false, false); err != nil {
		return err
	}
	if err := checkFileExist(desPath, true, false); err != nil {
		return err
	}
	if obj, err = handler.NewFileHandler(&model.Data{Path: desPath, InBin: false}).GetTarget(); err != nil {
		return err
	}
	if err = handler.SysCompare(srcPath, filepath.Join(model.Root, obj.Path)); err != nil {
		return err
	}
	return nil
}

func Delete(path string) error {
	if err := checkFileExist(path, true, false); err != nil {
		return err
	}
	if err := checkIsDir(path, true, false); err == nil {
		//is dir
		var files []model.Data
		if files, err = handler.NewFileHandler(&model.Data{Path: path, InBin: false}).GetAllInDir(); err != nil {
			return err
		}
		for i := range files {
			files[i].BinPath = path[len(filepath.Dir(path)):]
			if files[i].BinPath == "" {
				files[i].BinPath = "/"
			}
			files[i].InBin = true
			if err = handler.NewFileHandler(&files[i]).MoveToBin(); err != nil {
				return err
			}
		}
	}

	obj, err := handler.NewFileHandler(&model.Data{Path: path, InBin: false}).GetTarget()
	if err != nil {
		return err
	}
	obj.BinPath = "/"
	obj.InBin = true
	if father, err := handler.NewFileHandler(&model.Data{BinPath: "/", InBin: true}).GetTarget(); err == nil {
		obj.PID = father.ID
	} else {
		return err
	}
	if err = handler.NewFileHandler(obj).MoveToBin(); err != nil {
		return err
	}
	if err := handler.SysMove(filepath.Join(model.Root, path), filepath.Join(model.Bin, obj.BinPath)); err != nil {
		return err
	}
	return nil
}

func Backup(srcPath, desPath string) error {
	if err := checkFileExist(srcPath, false, false); err != nil {
		return err
	}
	if obj, err = handler.SysReadFileInfo(srcPath); err != nil {
		return err
	}
	obj.Path = filepath.Join(desPath, filepath.Base(srcPath))
	if err := handler.NewFileHandler(obj).Backup(); err != nil {
		return err
	}
	if err := handler.SysCopy(srcPath, filepath.Join(model.Root, obj.Path), uint32(os.ModePerm), time.Now()); err != nil {
		return err
	}

	return nil
}

func BackupWithKey(srcPath, desPath, key string) error {
	if err := checkFileExist(srcPath, false, false); err != nil {
		return err
	}
	//todo
	return nil
}

func Clean(path string) error {
	if err := checkFileExist(path, true, true); err != nil {
		return err
	}
	if err := checkIsDir(path, true, true); err == nil {
		//is dir
		var files []model.Data
		if files, err = handler.NewFileHandler(&model.Data{BinPath: path, InBin: true}).GetAllInDir(); err != nil {
			return err
		}
		for i := range files {
			if err = handler.NewFileHandler(&files[i]).Clean(); err != nil {
				return err
			}
		}
	}

	obj, err := handler.NewFileHandler(&model.Data{BinPath: path, InBin: true}).GetTarget()
	if err != nil {
		return err
	}
	if err = handler.NewFileHandler(obj).Clean(); err != nil {
		return err
	}
	if err := handler.SysClean(filepath.Join(model.Bin, path)); err != nil {
		return err
	}
	return nil
}

func Recycle(path string) error {
	if err := checkFileExist(path, true, true); err != nil {
		return err
	}
	if err := checkIsDir(path, true, true); err == nil {
		//is dir
		var files []model.Data
		if files, err = handler.NewFileHandler(&model.Data{BinPath: path, InBin: true}).GetAllInDir(); err != nil {
			return err
		}
		for i := range files {
			files[i].InBin = false
			if err = handler.NewFileHandler(&files[i]).Recycle(); err != nil {
				return err
			}
		}
	}

	obj, err := handler.NewFileHandler(&model.Data{BinPath: path, InBin: true}).GetTarget()
	if err != nil {
		return err
	}
	obj.InBin = false
	if father, err := handler.NewFileHandler(&model.Data{Path: filepath.Dir(obj.Path), InBin: false}).GetTarget(); err == nil {
		obj.PID = father.ID
	} else {
		return err
	}

	if err = handler.NewFileHandler(obj).Recycle(); err != nil {
		return err
	}
	if err := handler.SysMove(filepath.Join(model.Bin, path), filepath.Join(model.Root, obj.Path)); err != nil {
		return err
	}
	return nil

}
