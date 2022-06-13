package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"github.com/zhangzqs/exam-system/service"
	"strconv"
)

type Question[T any] struct {
	Type    string `json:"type"`
	Content T      `json:"content"`
}

func AddQuestion(c *gin.Context) {
	uid := GetUid(c)
	mp := make(map[string]any)
	err := c.BindJSON(&mp)
	if err != nil {
		RequestFormatError(c)
		return
	}

	var id int
	qType := mp["type"]
	qContent := mp["content"]
	switch qType {
	case "single":
		var q service.SingleQuestion
		err = mapstructure.Decode(qContent, &q)
		if err != nil {
			RequestFormatError(c, "单选题内容解析异常", err)
		}
		id, err = service.AddSingleQuestion(uid, &q)

	case "multiple":
		var q service.MultipleQuestion
		err = mapstructure.Decode(qContent, &q)
		if err != nil {
			RequestFormatError(c, "多选题内容解析异常", err)
		}
		id, err = service.AddMultipleQuestion(uid, &q)
	case "fill":
		var q service.FillQuestion
		err = mapstructure.Decode(qContent, &q)
		if err != nil {
			RequestFormatError(c, "填空题内容解析异常", err)
		}
		id, err = service.AddFillQuestion(uid, &q)
	case "judge":
		var q service.JudgeQuestion
		err = mapstructure.Decode(qContent, &q)
		if err != nil {
			RequestFormatError(c, "判断题内容解析异常", err)
		}
		id, err = service.AddJudgeQuestion(uid, &q)
	default:
		RequestContentError(c, "题目类型错误：", qType)
		return
	}
	if err != nil {
		DatabaseError(c, "题目存储异常", err)
		return
	}

	SuccessfulApiResponse(c, gin.H{
		"id": id,
	})
}
func DeleteQuestion(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RequestContentError(c, err.Error())
	}
	uid := GetUid(c)
	service.DeleteQuestion(uid, id)
	SuccessfulApiResponse(c)
}
func UpdateQuestion(c *gin.Context) {
	uid := GetUid(c)

	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RequestFormatError(c)
		return
	}
	mp := make(map[string]any)
	err = c.BindJSON(&mp)
	if err != nil {
		RequestFormatError(c)
		return
	}

	qType := mp["type"]
	qContent := mp["content"]

	switch qType {
	case "single":
		var q service.SingleQuestion
		err = mapstructure.Decode(qContent, &q)
		if err != nil {
			RequestFormatError(c, "单选题内容解析异常", err)
		}
		err = service.UpdateSingleQuestion(uid, id, &q)

	case "multiple":
		var q service.MultipleQuestion
		err = mapstructure.Decode(qContent, &q)
		if err != nil {
			RequestFormatError(c, "多选题内容解析异常", err)
		}
		err = service.UpdateMultipleQuestion(uid, id, &q)
	case "fill":
		var q service.FillQuestion
		err = mapstructure.Decode(qContent, &q)
		if err != nil {
			RequestFormatError(c, "填空题内容解析异常", err)
		}
		err = service.UpdateFillQuestion(uid, id, &q)
	case "judge":
		var q service.JudgeQuestion
		err = mapstructure.Decode(qContent, &q)
		if err != nil {
			RequestFormatError(c, "判断题内容解析异常", err)
		}
		err = service.UpdateJudgeQuestion(uid, id, &q)
	default:
		RequestContentError(c, "题目类型错误：", qType)
		return
	}
	if err != nil {
		DatabaseError(c, "题目存储异常", err)
		return
	}

	SuccessfulApiResponse(c)
}

func GetQuestion(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RequestContentError(c, err.Error())
	}
	uid := GetUid(c)

	service.GetQuestion(uid, id)
	SuccessfulApiResponse(c)
}

func GetUserQuestions(c *gin.Context) {
	uid := GetUid(c)

	service.GetUserQuestions(uid)
	SuccessfulApiResponse(c)
}
