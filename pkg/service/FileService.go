package service

import (
	"SE_Project/pkg/model"
	"errors"
	"log"
	"os"
)

func CheckIsDir(path string) error {
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

func ReadDir(path string) ([]model.ObjectPointer, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	files, err := f.ReadDir(0)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var result []model.ObjectPointer
	for _, file := range files {
		log.Println(file)
		info, err := file.Info()
		if err != nil {
			return nil, err
		}
		result = append(result, model.ObjectPointer{
			Name: file.Name(), 
			Type: file.Type().String(), 
			Size: uint64(info.Size())})
	}
	return result, nil
}
