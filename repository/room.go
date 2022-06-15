package repository

import (
	"github.com/zhangzqs/exam-system/global"
	"time"
)

type Student struct {
	Uid      int        `json:"uid"`
	EnterAt  *time.Time `json:"enterAt"`
	SubmitAt *time.Time `json:"submitAt"`
	Comment  *string    `json:"comment"`
	Score    *float64   `json:"score"`
}
type RoomEntity struct {
	RoomId      int       `json:"roomId"`
	PaperId     int       `json:"paperId"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	StudentList []Student `json:"studentList"`
}

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

func GetRoomsByUid(uid int) ([]RoomEntity, error) {
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
	var rooms []RoomEntity
	for cur.Next() {
		var r RoomEntity
		r.StudentList = []Student{}

		if err := cur.Scan(&r.RoomId, &r.PaperId, &r.StartTime, &r.EndTime); err != nil {
			return nil, err
		}
		rooms = append(rooms, r)
	}
	return rooms, nil
}
func GetRoom(roomId int) (*RoomEntity, error) {
	var r RoomEntity
	r.StudentList = []Student{}
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
		var stu Student
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

func UpdateEnterRoomTime(roomId int, studentId int) error {
	db := global.GetDatabase()
	if _, err := db.Exec(
		"UPDATE user_room "+
			"SET enter_at = CASE "+
			"    WHEN enter_at IS NULL THEN now() "+
			"    ELSE enter_at "+
			"END "+
			"WHERE rid = $1 "+
			"  AND uid = $2",
		roomId, studentId,
	); err != nil {
		return err
	}
	return nil
}
