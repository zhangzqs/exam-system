package repository

import (
	"errors"
	"github.com/zhangzqs/exam-system/global"
)

// UserExists 用户名是否存在
func UserExists(username string) bool {
	db := global.GetDatabase()
	cur, _ := db.Query(
		"SELECT * FROM users WHERE username=$1",
		username,
	)
	return cur.Next()
}

// UserValid 验证用户
func UserValid(username string, password string) (uid int, err error) {
	db := global.GetDatabase()
	cur, err := db.Query(
		"SELECT uid FROM users WHERE username=$1 AND password=$2",
		username,
		password,
	)
	if err != nil {
		return -1, err
	}
	if cur.Next() {
		err := cur.Scan(&uid)
		if err != nil {
			return -1, err
		}
		return uid, nil
	}
	return -1, errors.New("用户名或密码不存在")
}

// UserInsert 注册用户
func UserInsert(username string, password string) error {
	db := global.GetDatabase()
	_, err := db.Exec("INSERT INTO users(username,password) VALUES ($1,$2)", username, password)
	if err != nil {
		return err
	}
	return err
}
