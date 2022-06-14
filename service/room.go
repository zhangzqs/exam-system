package service

import (
	"github.com/zhangzqs/exam-system/entity"
	"github.com/zhangzqs/exam-system/repository"
)

// CreateRoom 创建一个考场
func CreateRoom(r *entity.RoomEntity) (rid int, err error) {
	rid, err = repository.InsertRoom(r.PaperId, r.StartTime, r.EndTime)
	if err != nil {
		return 0, err
	}
	r.RoomId = rid
	r.StudentList = []entity.Student{}

	return
}
