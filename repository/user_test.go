package repository

import (
	"fmt"
	"github.com/zhangzqs/exam-system/global"
	"testing"
)

func TestUserValid(t *testing.T) {
	global.InitConfig(&global.Config{
		Hostname: "localhost",
		Username: "postgres",
		Password: "123456",
		Port:     5432,
		DbName:   "postgres",
		SslMode:  "disable",
	})
	//err := UserInsert("zzq", "pwd")
	//log.Println(err)
	uid, err := UserValid("zzq", "pwd")
	fmt.Println(uid, err)
}
