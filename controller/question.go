package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/service"
	"strconv"
)

type addQuestionRequestBody struct {
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Options []string `json:"options"`

	// int | []int | []string | bool
	Answer interface{} `json:"answer"`
}
type addQuestionResponseBody struct {
	Id int `json:"id"`
}

func AddQuestion(c *gin.Context) {
	//uid := GetUid(c)
	var qr addQuestionRequestBody
	err := c.BindJSON(&qr)
	if err != nil {
		RequestFormatError(c)
		return
	}

	//switch qr.Type {
	//case "single":
	//	SuccessfulApiResponse(c, addQuestionResponseBody{
	//		Id: service.AddSingleQuestion(uid, &service.SingleQuestion{
	//			Title: qr.Title, Options: qr.Options, Answer: qr.Answer.(int),
	//		}),
	//	})
	//	return
	//case "multiple":
	//	SuccessfulApiResponse(c, addQuestionResponseBody{
	//		Id: service.AddMultipleQuestion(uid, qr.Title, qr.Options, qr.Answer.([]int)),
	//	})
	//	return
	//case "fill":
	//	SuccessfulApiResponse(c, addQuestionResponseBody{
	//		Id: service.AddFillQuestion(uid, qr.Title, qr.Answer.([]string)),
	//	})
	//	return
	//case "judge":
	//	SuccessfulApiResponse(c, addQuestionResponseBody{
	//		Id: service.AddJudgeQuestion(uid, qr.Title, qr.Answer.(bool)),
	//	})
	//	return
	//
	//default:
	//	RequestContentError(c, "题目类型错误："+qr.Type)
	//}
	RequestContentError(c, "题目类型错误："+qr.Type)
}
func DeleteQuestion(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RequestContentError(c, err.Error())
	}
	uid := GetUid(c)
	service.DeleteQuestion(uid, id)
	SuccessfulApiResponse(c, nil)
}

type updateQuestionRequestBody struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Options []string `json:"options"`

	// int | []int | []string | bool
	Answer interface{} `json:"answer"`
}

func UpdateQuestion(c *gin.Context) {
	uid := GetUid(c)

	var qr updateQuestionRequestBody
	err := c.BindJSON(&qr)
	if err != nil {
		RequestFormatError(c)
		return
	}

	switch qr.Type {
	case "single":
		service.UpdateSingleQuestion(uid, qr.Id, qr.Title, qr.Options, qr.Answer.(int))
		SuccessfulApiResponse(c, nil)
		return
	case "multiple":
		service.UpdateMultipleQuestion(uid, qr.Id, qr.Title, qr.Options, qr.Answer.([]int))
		SuccessfulApiResponse(c, nil)
		return
	case "fill":
		service.UpdateFillQuestion(uid, qr.Id, qr.Title, qr.Answer.([]string))
		SuccessfulApiResponse(c, nil)
		return
	case "judge":
		service.UpdateJudgeQuestion(uid, qr.Id, qr.Title, qr.Answer.(bool))
		SuccessfulApiResponse(c, nil)
		return

	default:
		RequestContentError(c, "题目类型错误："+qr.Type)
	}
}

func GetQuestion(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RequestContentError(c, err.Error())
	}
	uid := GetUid(c)

	service.GetQuestion(uid, id)
	SuccessfulApiResponse(c, nil)
}

func GetUserQuestions(c *gin.Context) {
	uid := GetUid(c)

	service.GetUserQuestions(uid)
	SuccessfulApiResponse(c, nil)
}
