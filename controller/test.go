package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/repository"
	"github.com/zhangzqs/exam-system/service"
	"strconv"
)

func GetCountScore(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	uid := GetUid(c)
	score, err := repository.CountScore(uid, roomId)
	if err != nil {
		DatabaseError(c, err)
		return
	}
	SuccessfulApiResponse(c, gin.H{
		"score": score,
	})
}

func GetScoreAndComment(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	comment, score, err := repository.GetCommentAndScore(roomId, uid)
	if err != nil {
		DatabaseError(c, err)
		return
	}
	SuccessfulApiResponse(c, gin.H{
		"score":   score,
		"comment": comment,
	})
}

func PutScoreAndComment(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}

	var req struct {
		Score   float64 `json:"score"`
		Comment string  `json:"comment"`
	}

	if err := c.BindJSON(&req); err != nil {
		RequestFormatError(c, err)
		return
	}

	if err := repository.UpdateCommentAndScore(roomId, uid, req.Comment, req.Score); err != nil {
		DatabaseError(c, err)
		return
	}
	SuccessfulApiResponse(c)
}

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
	roomId, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		RequestFormatError(c, err)
		return
	}
	uid := GetUid(c)
	comment, score, err := repository.GetCommentAndScore(roomId, uid)
	if err != nil {
		DatabaseError(c, err)
		return
	}
	SuccessfulApiResponse(c, gin.H{
		"score":   score,
		"comment": comment,
	})
}
