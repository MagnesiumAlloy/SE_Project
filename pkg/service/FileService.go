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

func SqlCheckIsDir(path string) error {
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

func SqlReadDir(path string) ([]model.Data, error) {
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

func CheckIsDir(dir *model.Data, readRoot bool) error {
	if readRoot {
		if err := handler.NewFileHandler(dir).CheckIsDir(); err != nil {
			return err
		}
		return nil
	} else {
		return SqlCheckIsDir(dir.Path + "/" + dir.Name)
	}

}

func ReadDir(dir *model.Data, readRoot bool) ([]model.Data, error) {
	if readRoot {

		result, err := handler.NewFileHandler(&model.Data{Path: dir.Path + "/" + dir.Name}).ReadDir()
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return SqlReadDir(dir.Path + "/" + dir.Name)
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
