package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/repository"
	"github.com/zhangzqs/exam-system/service"
	"strconv"
)

func CreateRoom(c *gin.Context) {
	var room repository.RoomEntity
	err := c.BindJSON(&room)
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	paperInfo, err := repository.GetPaperInfo(room.PaperId)
	if err != nil {
		DatabaseError(c, err)
		return
	}
	if paperInfo.CreatedBy != GetUid(c) {
		PermissionError(c, "您无权限使用试卷：", paperInfo.Pid)
		return
	}
	rid, err := service.CreateRoom(&room)
	if err != nil {
		DatabaseError(c, err)
		return
	}
	room.RoomId = rid
	SuccessfulApiResponse(c, room)
}

func GetUserRooms(c *gin.Context) {
	rooms, err := repository.GetRoomsByUid(GetUid(c))
	if err != nil {
		DatabaseError(c, err)
		return
	}
	SuccessfulApiResponse(c, gin.H{
		"rooms": rooms,
	})
}

func AddStudents(c *gin.Context) {
	var uids []int
	err := c.BindJSON(&uids)
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	for _, uid := range uids {
		err := repository.AddStudent(roomId, uid)
		if err != nil {
			DatabaseError(c, err)
			return
		}
	}
	SuccessfulApiResponse(c)
}

func GetRoom(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	room, err := repository.GetRoom(roomId)
	if err != nil {
		DatabaseError(c, err)
		return
	}
	SuccessfulApiResponse(c, room)
}
