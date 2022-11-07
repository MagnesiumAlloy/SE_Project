package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"log"
	"path/filepath"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func initDB() {
	//handler.GetDB().Debug().Raw("drop table data")
	handler.GetDB().AutoMigrate(&model.User{})
	handler.GetDB().AutoMigrate(&model.Data{})
	handler.GetDB().Exec("DELETE FROM data")
	//handler.GetDB().Raw("ALTER TABLE data ADD UNIQUE KEY(path, name);alter table data modify name varchar(256);")
	pwd, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	handler.GetDB().Create(&model.User{UserName: "user", Password: string(pwd), UserType: model.NormalUserType})
	pwd, _ = bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	handler.GetDB().Create(&model.User{UserName: "admin", Password: string(pwd), UserType: model.AdminType})

	res, _ := handler.ReadAllFileAndDir(model.Root)
	for _, x := range res {
		x.Path = x.Path[len(model.Root):]
		if x.Path == "" {
			x.Path = "/"
		} else {
			var f model.Data
			handler.GetDB().Where(&model.Data{Path: filepath.Dir(x.Path)}).First(&f)
			x.PID = f.ID
		}
		handler.GetDB().Create(&x)
	}
	res, _ = handler.ReadAllFileAndDir(model.Bin)
	for _, x := range res {
		x.Path = x.Path[len(model.Bin):]
		if x.Path == "" {
			x.Path = "/"
		} else {
			var f model.Data
			handler.GetDB().Where(&model.Data{Path: filepath.Dir(x.Path)}).First(&f)
			x.PID = f.ID
		}
		x.InBin = true
		x.BinPath = x.Path
		handler.GetDB().Create(&x)
	}
}

func TestBackup(t *testing.T) {
	model.Root = "/home/lush/Cloud_Backup"
	model.Bin = "/home/lush/Cloud_Bin"
	if err := Backup("/home/lush/tmp/dir", "/dir"); err != nil {
		log.Println(err.Error())
	}
}

func TestClean(t *testing.T) {

}

func TestDelete(t *testing.T) {
	model.Root = "/home/lush/Cloud_Backup"
	model.Bin = "/home/lush/Cloud_Bin"
	initDB()
	if err := Delete("/dir"); err != nil {
		log.Println(err.Error())
	}
}

func TestRecycle(t *testing.T) {
	model.Root = "/home/lush/Cloud_Backup"
	model.Bin = "/home/lush/Cloud_Bin"
	initDB()
	if err := Recycle("/dir/pl.txt"); err != nil {
		log.Println(err.Error())
	}
	if err := Recycle("/dir"); err != nil {
		log.Println(err.Error())
	}
	if err := Recycle("/main.go"); err != nil {
		log.Println(err.Error())
	}
}

func TestCompare(t *testing.T) {
	model.Root = "/home/lush/Cloud_Backup"
	model.Bin = "/home/lush/Cloud_Bin"
	initDB()
	if err := Compare("/home/lush/tmp/main.go", "/main.go"); err != nil {
		log.Println(err.Error())
	}
	if err := Compare("/home/lush/tmp/dir/nihao.go", "/main.go"); err != nil {
		log.Println(err.Error())
	}
	if err := Compare("/home/lush/tmp/main.go", "/dir"); err != nil {
		log.Println(err.Error())
	}
}

func TestRecover(t *testing.T) {
	model.Root = "/home/lush/Cloud_Backup"
	model.Bin = "/home/lush/Cloud_Bin"
	initDB()
	if err := Recover("/main.go", "/home/lush/tmp"); err != nil {
		log.Println(err.Error())
	}

	if err := Recover("/dir", "/home/lush/tmp"); err != nil {
		log.Println(err.Error())
	}
}

func TestBackupPackedData(t *testing.T) {
	BackupPackedData("/home/lush/SE_Project/web/html", "/")
}
func TestRecoverPackedData(t *testing.T) {
	RecoverPackedData("/home/lush/tmp/static.cloud", "/home/lush/tmp/newdir")
}
