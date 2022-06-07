package service

import (
	"errors"
	"github.com/zhangzqs/exam-system/global"
	"github.com/zhangzqs/exam-system/repository"
)

func Login(username string, password string) (string, error) {
	uid, err := repository.UserValid(username, password)
	if err != nil {
		return "", err
	}
	jwt := global.GetJwt()
	return jwt.GenerateToken(uid), nil
}

func Register(username string, password string) (string, error) {
	err := repository.UserInsert(username, password)
	if err != nil {
		if repository.UserExists(username) {
			return "", errors.New("用户已存在")
		}
		return "", err
	}
	return Login(username, password)
}
