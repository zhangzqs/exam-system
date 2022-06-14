package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/entity"
	"github.com/zhangzqs/exam-system/service"
)

func CreateRoom(c *gin.Context) {
	var room entity.RoomEntity
	err := c.BindJSON(&room)
	if err != nil {
		RequestFormatError(c, err)
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
