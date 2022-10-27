package validator

import "errors"

func CheckFileName(name string) error {
	if len(name) <= 0 || len(name) > 256 {
		return errors.New("invalid file name")
	}
	return nil
}

func CheckPath(path string) error {
	if len(path) > 0 && path[0] != '/' || len(path) > 256 {
		return errors.New("invalid path string")
	}
	return nil
}

func CheckNameAndPath(names, paths []string) error {
	for _, name := range names {
		if err := CheckFileName(name); err != nil {
			return err
		}
	}
	for _, path := range paths {
		if err := CheckPath(path); err != nil {
			return err
		}
	}
	return nil
}
