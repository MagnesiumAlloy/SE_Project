package util

import (
	"SE_Project/internal/model"
	"errors"
)

func GetTargetType(name string) string {
	res := model.Dir
	tmp := ""
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			break
		}
		if name[i] == '.' {
			res = tmp
			break
		}
		tmp = string(name[i]) + tmp
	}
	return res
}

func PrefixReplace(oldpre, newpre, s string) (string, error) {
	if len(oldpre) > len(s) {
		return s, errors.New("invalid prefix")
	}
	for i := range oldpre {
		if s[i] != oldpre[i] {
			return s, errors.New("invalid prefix")
		}
	}
	return newpre + s[len(oldpre):], nil
}

func PasswordPadding(pwd string) []byte {
	res := []byte(pwd)
	for i := 0; i < 16-len(pwd); i++ {
		res = append(res, byte(pwd[i%len(pwd)]+100))
	}
	return res
}
