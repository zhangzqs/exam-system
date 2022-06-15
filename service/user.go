package service

import (
	"errors"
	"github.com/zhangzqs/exam-system/repository"
)

func Login(username string, password string) (int, error) {
	uid, err := repository.UserValid(username, password)
	if err != nil {
		return -1, err
	}
	return uid, nil
}

func Register(username string, password string) (int, error) {
	err := repository.UserInsert(username, password)
	if err != nil {
		if repository.UserExists(username) {
			return -1, errors.New("用户已存在")
		}
		return -1, err
	}
	return Login(username, password)
}
