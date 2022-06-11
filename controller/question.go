package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/service"
)

type addQuestionRequestBody struct {
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
	Answer  interface {
		int | []int | []string | bool
	} `json:"answer"`
}
type addQuestionResponseBody struct {
	Id int `json:"id"`
}

func AddQuestion(c *gin.Context) {
	var qr addQuestionRequestBody
	err := c.BindJSON(&qr)
	if err != nil {
		RequestFormatError(c)
		return
	}

	switch qr.Type {
	case "single":
		SuccessfulApiResponse(c, addQuestionResponseBody{
			Id: service.AddSingleQuestion(qr.Title, qr.Options, qr.Answer.(int)),
		})
		return
	case "multiple":
		SuccessfulApiResponse(c, addQuestionResponseBody{
			Id: service.AddMultipleQuestion(qr.Title, qr.Options, qr.Answer.([]int)),
		})
		return
	case "fill":
		SuccessfulApiResponse(c, addQuestionResponseBody{
			Id: service.AddFillQuestion(qr.Title, qr.Answer.([]string)),
		})
		return
	case "judge":
		SuccessfulApiResponse(c, addQuestionResponseBody{
			Id: service.AddJudgeQuestion(qr.Title, qr.Answer.(bool)),
		})
		return

	default:
		RequestContentError(c, "题目类型错误："+qr.Type)
	}
}

type questionIdRequestBody struct {
	Id int `json:"id"`
}

func DeleteQuestion(c *gin.Context) {
	var qr questionIdRequestBody
	err := c.BindJSON(&qr)
	if err != nil {
		RequestFormatError(c)
		return
	}
	service.DeleteQuestion(qr.Id)
	SuccessfulApiResponse(c, nil)
}

type updateQuestionRequestBody struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
	Answer  interface {
		int | []int | []string | bool
	} `json:"answer"`
}

func UpdateQuestion(c *gin.Context) {
	var qr updateQuestionRequestBody
	err := c.BindJSON(&qr)
	if err != nil {
		RequestFormatError(c)
		return
	}

	switch qr.Type {
	case "single":
		service.UpdateSingleQuestion(qr.Id, qr.Title, qr.Options, qr.Answer.(int))
		SuccessfulApiResponse(c, nil)
		return
	case "multiple":
		service.UpdateMultipleQuestion(qr.Id, qr.Title, qr.Options, qr.Answer.([]int))
		SuccessfulApiResponse(c, nil)
		return
	case "fill":
		service.UpdateFillQuestion(qr.Id, qr.Title, qr.Answer.([]string))
		SuccessfulApiResponse(c, nil)
		return
	case "judge":
		service.UpdateJudgeQuestion(qr.Id, qr.Title, qr.Answer.(bool))
		SuccessfulApiResponse(c, nil)
		return

	default:
		RequestContentError(c, "题目类型错误："+qr.Type)
	}
}

func GetQuestion(c *gin.Context) {
	var qr questionIdRequestBody
	err := c.BindJSON(&qr)
	if err != nil {
		RequestFormatError(c)
		return
	}
	service.GetQuestion(qr.Id)
	SuccessfulApiResponse(c, nil)
}

func GetUserQuestions(c *gin.Context) {
	uid := 12
	service.GetUserQuestions(uid)
	SuccessfulApiResponse(c, nil)
}
