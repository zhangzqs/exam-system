package repository

import (
	"github.com/zhangzqs/exam-system/entity"
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

func GetRoomsByUid(uid int) ([]entity.RoomEntity, error) {
	db := global.GetDatabase()
	cur, err := db.Query(
		"SELECT rid, rooms.pid, start_time, end_time "+
			"FROM rooms, "+
			"     (SELECT DISTINCT pid "+
			"      FROM papers "+
			"      WHERE created_by = $1) as X "+
			"WHERE rooms.pid = X.pid", uid)
	if err != nil {
		return nil, err
	}
	var rooms []entity.RoomEntity
	for cur.Next() {
		var r entity.RoomEntity
		r.StudentList = []entity.Student{}

		if err := cur.Scan(&r.RoomId, &r.PaperId, &r.StartTime, &r.EndTime); err != nil {
			return nil, err
		}
		rooms = append(rooms, r)
	}
	return rooms, nil
}
func GetRoom(roomId int) (*entity.RoomEntity, error) {
	var r entity.RoomEntity
	r.StudentList = []entity.Student{}
	db := global.GetDatabase()
	if err := db.QueryRow(
		"SELECT rid,pid,start_time,end_time "+
			"FROM rooms "+
			"WHERE rid=$1",
		roomId,
	).Scan(&r.RoomId, &r.PaperId, &r.StartTime, &r.EndTime); err != nil {
		return nil, err
	}

	cur, err := db.Query(
		"SELECT uid,enter_at,submit_at,comment,score "+
			"FROM user_room "+
			"WHERE rid=$1", roomId)
	if err != nil {
		return nil, err
	}
	for cur.Next() {
		var stu entity.Student
		if err := cur.Scan(&stu.Uid, &stu.EnterAt, &stu.SubmitAt, &stu.Comment, &stu.Score); err != nil {
			return nil, err
		}
		r.StudentList = append(r.StudentList, stu)
	}
	return &r, nil
}

func AddStudent(roomId int, studentId int) error {
	db := global.GetDatabase()
	if _, err := db.Exec(
		"INSERT INTO user_room(uid,rid) VALUES ($1,$2)",
		studentId, roomId,
	); err != nil {
		return err
	}
	return nil
}
