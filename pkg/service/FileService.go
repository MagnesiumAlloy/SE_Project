package service

import (
	"SE_Project/internal/model"
	"SE_Project/internal/util"
	"SE_Project/pkg/handler"
	"errors"
	"fmt"
	"os"
	"os/user"
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
				if userID != 0 {
					path = filepath.Join(path, fmt.Sprint(userID))
				}
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
				if userID != 0 {
					path = filepath.Join(path, fmt.Sprint(userID))
				}
				println(path)
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
	if desPath == "/" {
		desPath = filepath.Join(desPath, fmt.Sprint(userID))
	}

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
	binPath := "/" + fmt.Sprint(userID)
	if err := checkFileExist(path, true, false, userID); err != nil {
		return err
	}
	if err := checkIsDir(path, true, false); err == nil {
		var PID uint
		if father, err := handler.NewFileHandler(&model.Data{BinPath: binPath, InBin: true, UserId: userID}).GetTarget(); err == nil {
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
			if i == 0 {
				files[i].PID = PID
			}
			files[i].BinPath = filepath.Join(binPath, files[i].Path[len(filepath.Dir(path)):])

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
		obj.BinPath = filepath.Join(binPath, obj.Name)
		obj.InBin = true
		if father, err := handler.NewFileHandler(&model.Data{BinPath: binPath, InBin: true, UserId: userID}).GetTarget(); err == nil {
			obj.PID = father.ID
		} else {
			return err
		}
		if err = handler.NewFileHandler(obj).MoveToBin(); err != nil {
			return err
		}

	}
	if err := handler.SysMove(filepath.Join(model.Root, path), filepath.Join(model.Bin, binPath, filepath.Base(path))); err != nil {
		return err
	}

	return nil
}

func BackupPackedData(srcPath, desPath string) error {
	var list []model.Data
	if list, err = handler.ReadAllFileAndDir(srcPath); err != nil {
		return err
	}
	/*
		for i, item := range list {
			list[i].Path = item.Path[len(filepath.Dir(srcPath)):]
			println(item.Type, item.Path, item.Size, item.ModTime.String())
		}
	*/
	desPath = filepath.Join(model.Root, desPath) + model.CloudTempType
	if obj, err = handler.SysWritePackedFile(filepath.Dir(srcPath), list, desPath); err != nil {
		return err
	}
	if err := util.Compress(desPath); err != nil {
		return err
	}
	if err := os.Remove(desPath); err != nil {
		return err
	}
	return nil
}

func RecoverPackedData(srcPath, desPath string) error {

	srcPath = filepath.Join(model.Root, srcPath)
	if err := util.Decompress(srcPath); err != nil {
		return err
	}

	_, err := handler.SysRecoverPackedFile(srcPath+model.CloudTempType, desPath)
	if err != nil {
		return err
	}
	if err := handler.SysClean(srcPath + model.CloudTempType); err != nil {
		return err
	}
	/*
		for _, item := range list {
			println(item.Type, item.Path, item.Size, item.ModTime.String())
		}
	*/
	return nil
}

func BackupData(srcPath, desPath string, userID uint) error {

	var list []model.Data
	if list, err = handler.ReadAllFileAndDir(srcPath); err != nil {
		return err
	}
	for _, item := range list {
		item.Path = filepath.Join(desPath, item.Path)
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
	if desPath == "/" {
		desPath = filepath.Join(desPath, fmt.Sprint(userID))
	}

	if err := checkFileExist(srcPath, false, false, userID); err != nil {
		return err
	}
	if pack {
		path := filepath.Join(desPath, filepath.Base(srcPath)) + model.CloudBackupType
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
		obj.UserId = userID
		obj.Path = obj.Path[len(model.Root):]
		var f *model.Data
		if f, err = handler.NewFileHandler(&model.Data{Path: filepath.Dir(obj.Path), UserId: userID}).GetTarget(); err != nil {
			return err
		}
		obj.PID = f.ID
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

//文件恢复
func Recover(srcPath, desPath, key string, userID uint) error {
	//检查待恢复文件存在
	if err = checkFileExist(srcPath, true, false, userID); err != nil {
		return err
	}
	//获取文件
	if obj, err = handler.NewFileHandler(&model.Data{Path: srcPath, InBin: false, UserId: userID}).GetTarget(); err != nil {
		return err
	}
	//检查是否加密
	if obj.Encrypted {
		//比较密码
		if err := bcrypt.CompareHashAndPassword([]byte(obj.Key), []byte(key)); err != nil {
			return err
		}
		//文件解密
		if err := handler.SysDecryptFile(filepath.Join(model.Root, srcPath), filepath.Join(model.Root, srcPath)+model.CloudTempType, key); err != nil {
			return err
		}
		//文件解包
		if err := RecoverPackedData(srcPath+model.CloudTempType, desPath); err != nil {
			return err
		}
		//清除临时文件
		return handler.SysClean(filepath.Join(model.Root, srcPath) + model.CloudTempType)
	}
	//检查是否打包存储
	if filepath.Ext(srcPath) == model.CloudBackupType {
		return RecoverPackedData(srcPath, desPath)
	}
	//恢复文件
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

func InitFileSys() {
	user, _ := user.Current()
	model.Bin = user.HomeDir + "/Cloud_Bin"
	model.Root = user.HomeDir + "/Cloud_Backup"
	//create if not exist

	if err := handler.SysCheckFileExist(model.Bin); err == nil {
		handler.SysClean(model.Bin)
	}
	os.Mkdir(model.Bin, os.ModePerm)

	if err := handler.SysCheckFileExist(model.Root); err == nil {
		handler.SysClean(model.Root)
	}
	os.Mkdir(model.Root, os.ModePerm)
	base := filepath.Join(filepath.Dir(model.Root), "/user_files")

	if err := handler.SysCheckFileExist(base); err == nil {
		handler.SysClean(base)
	}
	os.Mkdir(base, os.ModePerm)
	if err := handler.SysCopy("./web/static", base, uint32(os.ModePerm), time.Now()); err != nil {
		return
	}

}

func InitDB() {
	/*
		command := `./init.sh`
		cmd := exec.Command("/bin/bash", "-c", command)

		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
		fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	*/
	//handler.GetDB().Debug().Raw("drop table data")
	handler.GetDB().Migrator().DropTable(&model.User{})
	handler.GetDB().Migrator().DropTable(&model.Data{})
	handler.GetDB().AutoMigrate(&model.User{})
	handler.GetDB().AutoMigrate(&model.Data{})
	//handler.GetDB().Raw("ALTER TABLE data ADD UNIQUE KEY(path, name);alter table data modify name varchar(256);")

	res, _ := handler.ReadAllFileAndDir(model.Root)
	for _, x := range res {
		x.Path = x.Path[len("/"+filepath.Base(model.Root)):]
		if x.Path == "" {
			x.Path = "/"
		} else {
			var f model.Data
			handler.GetDB().Where(&model.Data{Path: filepath.Dir(x.Path)}).First(&f)
			x.PID = f.ID
		}
		x.InBin = false
		handler.GetDB().Create(&x)
	}
	res, _ = handler.ReadAllFileAndDir(model.Bin)
	for _, x := range res {
		x.Path = x.Path[len("/"+filepath.Base(model.Bin)):]
		if x.Path == "" {
			x.Path = "/"
		} else {
			var f model.Data
			handler.GetDB().Where(&model.Data{BinPath: filepath.Dir(x.Path)}).First(&f)
			x.PID = f.ID
		}
		x.InBin = true
		x.BinPath = x.Path
		handler.GetDB().Create(&x)
	}

	//service.Register("user", "123")
	//service.Register("admin", "admin")
}
