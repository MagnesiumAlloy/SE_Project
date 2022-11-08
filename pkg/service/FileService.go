package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"errors"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func ReadDir(path string, ID *int, isRoot bool, inBin bool, userID uint) ([]model.Data, error) {
	var res []model.Data
	if err := checkIsDir(path, isRoot, inBin); err != nil {
		return nil, err
	}
	if isRoot {
		if inBin {
			if path == "/" {
				if obj, err := handler.NewFileHandler(&model.Data{BinPath: path, InBin: inBin, UserId: userID}).GetTarget(); err != nil {
					return nil, err
				} else {
					*ID = int(obj.ID)
				}
			}
			if res, err = handler.NewFileHandler(&model.Data{PID: uint(*ID), InBin: inBin, UserId: userID}).ReadDir(); err != nil {
				return nil, err
			}
		} else {
			if path == "/" {
				if obj, err := handler.NewFileHandler(&model.Data{Path: path, InBin: inBin, UserId: userID}).GetTarget(); err != nil {
					return nil, err
				} else {
					*ID = int(obj.ID)
				}
			}
			if res, err = handler.NewFileHandler(&model.Data{PID: uint(*ID), InBin: inBin, UserId: userID}).ReadDir(); err != nil {
				return nil, err
			}
		}
		return res, nil
	} else {
		return handler.SysReadDir(path)
	}
}

func checkFileExist(path string, isRoot bool, inBin bool, userID uint) error {
	if isRoot {
		if inBin {
			if err := handler.NewFileHandler(&model.Data{BinPath: path, InBin: inBin, UserId: userID}).CheckTargetExist(); err != nil {
				return err
			}
		} else {
			if err := handler.NewFileHandler(&model.Data{Path: path, InBin: inBin, UserId: userID}).CheckTargetExist(); err != nil {
				return err
			}
		}

	} else {
		return handler.SysCheckFileExist(path)
	}
	return nil
}

func Compare(srcPath, desPath string, userID uint) error {
	if err := checkFileExist(srcPath, false, false, userID); err != nil {
		return err
	}
	if err := handler.SysCheckIsDir(srcPath); err == nil {
		return errors.New("source file is not a file")
	}
	if err := checkFileExist(desPath, true, false, userID); err != nil {
		return err
	}
	if err := handler.NewFileHandler(&model.Data{Path: desPath, InBin: false, Type: model.Dir, UserId: userID}).CheckTargetExist(); err == nil {
		return errors.New("des file is not a file")
	}
	if obj, err = handler.NewFileHandler(&model.Data{Path: desPath, InBin: false, UserId: userID}).GetTarget(); err != nil {
		return err
	}
	if err = handler.SysCompare(srcPath, filepath.Join(model.Root, obj.Path)); err != nil {
		return err
	}
	return nil
}

func Delete(path string, userID uint) error {
	if err := checkFileExist(path, true, false, userID); err != nil {
		return err
	}
	if err := checkIsDir(path, true, false); err == nil {
		var PID uint
		if father, err := handler.NewFileHandler(&model.Data{BinPath: "/", InBin: true, UserId: userID}).GetTarget(); err == nil {
			PID = father.ID
		} else {
			return err
		}
		//is dir
		var files []model.Data
		if files, err = handler.NewFileHandler(&model.Data{Path: path, InBin: false, UserId: userID}).GetAllInDir(); err != nil {
			return err
		}
		for i := range files {
			files[i].BinPath = "/" + path[len(filepath.Dir(path)):]
			if i == 0 {
				files[i].PID = PID
			}
			if files[i].BinPath == "" {
				files[i].BinPath = "/" + files[i].Name
			}
			files[i].InBin = true
			if err = handler.NewFileHandler(&files[i]).MoveToBin(); err != nil {
				return err
			}
		}
	} else {
		obj, err := handler.NewFileHandler(&model.Data{Path: path, InBin: false, UserId: userID}).GetTarget()
		if err != nil {
			return err
		}
		obj.BinPath = "/" + obj.Name
		obj.InBin = true
		if father, err := handler.NewFileHandler(&model.Data{BinPath: "/", InBin: true, UserId: userID}).GetTarget(); err == nil {
			obj.PID = father.ID
		} else {
			return err
		}
		if err = handler.NewFileHandler(obj).MoveToBin(); err != nil {
			return err
		}

	}
	if err := handler.SysMove(filepath.Join(model.Root, path), filepath.Join(model.Bin, filepath.Base(path))); err != nil {
		return err
	}

	return nil
}

func BackupPackedData(srcPath, desPath string) error {
	var list []model.Data
	if list, err = handler.ReadAllFileAndDir(srcPath); err != nil {
		return err
	}
	for i, item := range list {
		list[i].Path = item.Path[len(filepath.Dir(srcPath)):]
		println(item.Type, item.Path, item.Size, item.ModTime.String())
	}
	if obj, err = handler.SysWritePackedFile(filepath.Dir(srcPath), list, filepath.Join(model.Root, desPath)); err != nil {
		return err
	}
	return nil
}

func RecoverPackedData(srcPath, desPath string) error {
	list, err := handler.SysRecoverPackedFile(filepath.Join(model.Root, srcPath), desPath)
	if err != nil {
		return err
	}
	for _, item := range list {
		println(item.Type, item.Path, item.Size, item.ModTime.String())
	}
	return nil
}

func BackupData(srcPath, desPath string, userID uint) error {

	var list []model.Data
	if list, err = handler.ReadAllFileAndDir(srcPath); err != nil {
		return err
	}
	for _, item := range list {
		item.Path = filepath.Join(desPath, item.Path[len(filepath.Dir(srcPath)):])
		var f *model.Data
		if f, err = handler.NewFileHandler(&model.Data{Path: filepath.Dir(item.Path), UserId: userID}).GetTarget(); err != nil {
			return err
		}
		item.PID = f.ID
		item.UserId = userID
		if err := handler.NewFileHandler(&item).Backup(); err != nil {
			return err
		}
	}
	if err := handler.SysCopy(srcPath, filepath.Join(model.Root, desPath, filepath.Base(srcPath)), uint32(os.ModePerm), time.Now()); err != nil {
		return err
	}

	return nil
}

func RecoverData(srcPath, desPath string, obj *model.Data) error {
	if err = handler.SysCopy(filepath.Join(model.Root, obj.Path), filepath.Join(desPath, filepath.Base(srcPath)), obj.Perm, obj.ModTime); err != nil {
		return err
	}
	return nil
}

func Backup(srcPath, desPath, key string, userID uint, pack bool, encrypt bool) error {
	if err := checkFileExist(srcPath, false, false, userID); err != nil {
		return err
	}
	if pack {
		path := filepath.Join(desPath, filepath.Dir(srcPath)) + model.CloudBackupType
		if err := BackupPackedData(srcPath, path); err != nil {
			return err
		}
		if encrypt {
			if err := handler.SysEncryptFile(filepath.Join(model.Root, path), key); err != nil {
				return err
			}
		}

		//sql
		obj, err = handler.SysReadFileInfo(filepath.Join(model.Root, path))
		if err != nil {
			return err
		}
		if encrypt {
			var encryptKey []byte
			if encryptKey, err = bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost); err != nil {
				return err
			}
			obj.Key = string(encryptKey)
			obj.Encrypted = true
		}
		if err := handler.NewFileHandler(obj).Backup(); err != nil {
			return err
		}
	} else {
		if err := BackupData(srcPath, desPath, userID); err != nil {
			return err
		}
	}
	return nil
}
func Recover(srcPath, desPath, key string, userID uint) error {
	if err = checkFileExist(srcPath, true, false, userID); err != nil {
		return err
	}
	if obj, err = handler.NewFileHandler(&model.Data{Path: srcPath, InBin: false, UserId: userID}).GetTarget(); err != nil {
		return err
	}
	if obj.Encrypted {
		if err := bcrypt.CompareHashAndPassword([]byte(obj.Key), []byte(key)); err != nil {
			return err
		}
		if err := handler.SysDecryptFile(filepath.Join(model.Root, srcPath), filepath.Join(model.Root, srcPath)+model.CloudTempType, key); err != nil {
			return err
		}
		if err := RecoverPackedData(filepath.Join(model.Root, srcPath)+model.CloudTempType, desPath); err != nil {
			return err
		}
		return handler.SysClean(filepath.Join(model.Root, srcPath) + model.CloudTempType)
	}
	if filepath.Ext(srcPath) == model.CloudBackupType {
		return RecoverPackedData(srcPath, desPath)
	}
	return RecoverData(srcPath, desPath, obj)
}

func Clean(path string, userID uint) error {
	if err := checkFileExist(path, true, true, userID); err != nil {
		return err
	}
	if err := checkIsDir(path, true, true); err == nil {
		//is dir
		var files []model.Data
		if files, err = handler.NewFileHandler(&model.Data{BinPath: path, InBin: true, UserId: userID}).GetAllInDir(); err != nil {
			return err
		}
		for i := range files {
			if err = handler.NewFileHandler(&files[i]).Clean(); err != nil {
				return err
			}
		}
	} else {
		obj, err := handler.NewFileHandler(&model.Data{BinPath: path, InBin: true, UserId: userID}).GetTarget()
		if err != nil {
			return err
		}
		if err = handler.NewFileHandler(obj).Clean(); err != nil {
			return err
		}
	}

	if err := handler.SysClean(filepath.Join(model.Bin, path)); err != nil {
		return err
	}
	return nil
}

func Recycle(path string, userID uint) error {
	if err := checkFileExist(path, true, true, userID); err != nil {
		return err
	}
	if err := checkIsDir(path, true, true); err == nil {
		obj, err = handler.NewFileHandler(&model.Data{BinPath: path, InBin: true, UserId: userID}).GetTarget()
		if err != nil {
			return err
		}
		var PID uint
		if father, err := handler.NewFileHandler(&model.Data{Path: filepath.Dir(obj.Path), InBin: false, UserId: userID}).GetTarget(); err == nil {
			PID = father.ID
		} else {
			return err
		}
		//is dir
		var files []model.Data
		if files, err = handler.NewFileHandler(&model.Data{BinPath: path, InBin: true, UserId: userID}).GetAllInDir(); err != nil {
			return err
		}
		for i := range files {
			if i == 0 {
				files[i].PID = PID
			}
			files[i].InBin = false
			if err = handler.NewFileHandler(&files[i]).Recycle(); err != nil {
				return err
			}
		}
	} else {
		obj, err = handler.NewFileHandler(&model.Data{BinPath: path, InBin: true, UserId: userID}).GetTarget()
		if err != nil {
			return err
		}
		obj.InBin = false
		if father, err := handler.NewFileHandler(&model.Data{Path: filepath.Dir(obj.Path), InBin: false, UserId: userID}).GetTarget(); err == nil {
			obj.PID = father.ID
		} else {
			return err
		}

		if err = handler.NewFileHandler(obj).Recycle(); err != nil {
			return err
		}

	}
	if err := handler.SysMove(filepath.Join(model.Bin, path), filepath.Join(model.Root, obj.Path)); err != nil {
		return err
	}

	return nil

}
