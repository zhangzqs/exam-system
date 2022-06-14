package repository

import (
	"github.com/zhangzqs/exam-system/global"
	"time"
)

func InsertRoom(pid int, startTime time.Time, endTime time.Time) (rid int, err error) {
	db := global.GetDatabase()
	if err = db.QueryRow(
		"INSERT INTO rooms "+
			"(pid,start_time,end_time) "+
			"VALUES ($1, $2, $3) "+
			"RETURNING rid",
		pid, startTime, endTime,
	).Scan(&rid); err != nil {
		return
	}
	return
}
