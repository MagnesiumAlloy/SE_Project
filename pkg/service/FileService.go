package service

import (
	"SE_Project/pkg/handler"
	"SE_Project/pkg/model"
	"SE_Project/pkg/util"
	"errors"
	"log"
	"os"
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
			Path: path})

	}
	return result, nil
}

func CheckIsDir(path string) error {
	obj := &model.Data{Path: path, Type: model.Dir}
	if err := handler.NewFileHandler(obj).CheckIsDir(); err != nil {
		return err
	}
	return nil
}

func ReadDir(path string) ([]model.Data, error) {
	obj := &model.Data{Path: path}
	result, err := handler.NewFileHandler(obj).ReadDir()
	if err != nil {
		return nil, err
	}
	return result, nil
}
