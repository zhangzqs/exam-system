package global

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func GetDatabase() *sql.DB {
	dbOnce.Do(func() {
		log.Println("加载数据库连接")
		conf := GetConfig()
		connStr := fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
			conf.Db.Username,
			conf.Db.Password,
			conf.Db.Hostname,
			conf.Db.Port,
			conf.Db.DbName,
			conf.Db.SslMode,
		)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalln(err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
	})
	return db
}
