package service

import (
	"encoding/json"
	"github.com/zhangzqs/exam-system/repository"
)

// EnterRoom 进入考场
func EnterRoom(roomId int, uid int) (room *repository.RoomEntity, paper *PaperContent, err error) {

	room, err = repository.GetRoom(roomId)
	if err != nil {
		return
	}
	paper, err = GetPaper(uid, room.PaperId, false)
	if err != nil {
		return
	}

	err = repository.UpdateEnterRoomTime(roomId, uid)
	if err != nil {
		return
	}

	return
}

func SubmitAnswer(uid int, rid int, qid int, answer any) error {
	ansJson, err := json.Marshal(answer)
	if err != nil {
		return err
	}
	err = repository.SubmitAnswer(uid, rid, qid, string(ansJson))
	if err != nil {
		return err
	}
	return nil
}
