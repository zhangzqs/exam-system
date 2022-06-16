package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/repository"
	"github.com/zhangzqs/exam-system/service"
	"strconv"
)

func EnterRoom(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	room, paper, err := service.EnterRoom(roomId, GetUid(c))
	if err != nil {
		return
	}

	SuccessfulApiResponse(c, gin.H{
		"room": gin.H{
			"startTime": room.StartTime,
			"endTime":   room.EndTime,
		},
		"paper": paper,
	})
}

func SubmitAnswer(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	var request struct {
		UserAnswers map[string]any `json:"userAnswers"`
	}
	err = c.BindJSON(&request)
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	for k, v := range request.UserAnswers {
		qid, err := strconv.Atoi(k)
		if err != nil {
			RequestFormatError(c, err)
			return
		}
		err = service.SubmitAnswer(GetUid(c), roomId, qid, v)
		if err != nil {
			DatabaseError(c, err)
			return
		}
	}
	SuccessfulApiResponse(c)
}

func SubmitPaper(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	err = repository.UpdateSubmitTime(roomId, GetUid(c))
	if err != nil {
		DatabaseError(c, err)
	}
	SuccessfulApiResponse(c)
}
func GetTestResult(c *gin.Context) {

}
