package service

import (
	"github.com/zhangzqs/exam-system/repository"
)

// CreateRoom 创建一个考场
func CreateRoom(r *repository.RoomEntity) (rid int, err error) {
	rid, err = repository.InsertRoom(r.PaperId, r.StartTime, r.EndTime)
	if err != nil {
		return 0, err
	}
	r.RoomId = rid
	r.StudentList = []repository.Student{}

	return
}

// EnterRoom 进入考场
func EnterRoom(roomId int, uid int) (room *repository.RoomEntity, paper *PaperContent, err error) {

	room, err = repository.GetRoom(roomId)
	if err != nil {
		return
	}
	paper, err = GetPaper(uid, room.PaperId)
	if err != nil {
		return
	}

	err = repository.UpdateEnterRoomTime(roomId, uid)
	if err != nil {
		return
	}

	return
}
