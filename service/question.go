package service

import (
	"encoding/json"
	"github.com/zhangzqs/exam-system/repository"
	"strconv"
)

type QuestionType interface {
	SingleQuestion | MultipleQuestion | FillQuestion | JudgeQuestion
}

type SingleQuestion struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Answer  *int     `json:"answer"`
}

func AddSingleQuestion(operatorUid int, question *SingleQuestion) (int, error) {
	jsonStr, _ := json.Marshal(question.Options)
	answerIdStr, _ := json.Marshal(*question.Answer)
	qid, err := repository.AddQuestion(&repository.Question{
		CreatedBy: operatorUid,
		Title:     question.Title,
		Type:      "single",
		Options:   string(jsonStr),
		Answer:    string(answerIdStr),
	})
	if err != nil {
		return 0, err
	}
	return qid, nil
}

type MultipleQuestion struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Answer  []int    `json:"answer"`
}

func AddMultipleQuestion(operatorUid int, question *MultipleQuestion) (int, error) {
	jsonOptionsStr, _ := json.Marshal(question.Options)
	jsonAnswerStr, _ := json.Marshal(question.Answer)
	qid, err := repository.AddQuestion(&repository.Question{
		CreatedBy: operatorUid,
		Title:     question.Title,
		Type:      "multiple",
		Options:   string(jsonOptionsStr),
		Answer:    string(jsonAnswerStr),
	})
	if err != nil {
		return 0, err
	}
	return qid, nil
}

type FillQuestion struct {
	Title  string   `json:"title"`
	Answer []string `json:"answer"`
}

func AddFillQuestion(operatorUid int, question *FillQuestion) (int, error) {
	jsonAnswerStr, _ := json.Marshal(question.Answer)

	qid, err := repository.AddQuestion(&repository.Question{
		CreatedBy: operatorUid,
		Title:     question.Title,
		Type:      "fill",
		Options:   "[]",
		Answer:    string(jsonAnswerStr),
	})

	if err != nil {
		return 0, err
	}
	return qid, nil
}

type JudgeQuestion struct {
	Title  string `json:"title"`
	Answer *bool  `json:"answer"`
}

func AddJudgeQuestion(operatorUid int, question *JudgeQuestion) (int, error) {
	jsonAnswerStr, _ := json.Marshal(question.Answer)

	qid, err := repository.AddQuestion(&repository.Question{
		CreatedBy: operatorUid,
		Title:     question.Title,
		Type:      "judge",
		Options:   "[]",
		Answer:    string(jsonAnswerStr),
	})

	if err != nil {
		return 0, err
	}
	return qid, nil
}

func DeleteQuestion(operatorUid int, id int) bool {
	Todo()
	return true
}

func UpdateSingleQuestion(operatorUid int, qid int, question *SingleQuestion) error {
	Todo()
	return nil
}
func UpdateMultipleQuestion(operatorUid int, qid int, question *MultipleQuestion) error {
	Todo()
	return nil

}
func UpdateFillQuestion(operatorUid int, qid int, question *FillQuestion) error {
	Todo()
	return nil

}
func UpdateJudgeQuestion(operatorUid int, qid int, question *JudgeQuestion) error {
	Todo()
	return nil

}

// ?????????????????????????????????????????????
func qTabEntity2Single(question *repository.Question, containsAnswer bool) *SingleQuestion {
	var options []string
	_ = json.Unmarshal([]byte(question.Options), &options)
	ans, _ := strconv.Atoi(question.Answer)
	pans := &ans
	if !containsAnswer {
		pans = nil
	}
	return &SingleQuestion{
		Title:   question.Title,
		Options: options,
		Answer:  pans,
	}
}

// ?????????????????????????????????????????????
func qTabEntity2Multiple(question *repository.Question, containsAnswer bool) *MultipleQuestion {
	var options []string
	var answer []int
	_ = json.Unmarshal([]byte(question.Options), &options)
	_ = json.Unmarshal([]byte(question.Answer), &answer)
	if !containsAnswer {
		answer = nil
	}
	return &MultipleQuestion{
		Title:   question.Title,
		Options: options,
		Answer:  answer,
	}
}

// ??????????????????????????????????????????
func qTabEntity2Fill(question *repository.Question, containsAnswer bool) *FillQuestion {
	var answer []string
	_ = json.Unmarshal([]byte(question.Answer), &answer)
	if !containsAnswer {
		answer = nil
	}
	return &FillQuestion{
		Title:  question.Title,
		Answer: answer,
	}
}

// ??????????????????????????????????????????
func qTabEntity2Judge(question *repository.Question, containsAnswer bool) *JudgeQuestion {
	answer, _ := strconv.ParseBool(question.Answer)
	pans := &answer
	if !containsAnswer {
		pans = nil
	}
	return &JudgeQuestion{
		Title:  question.Title,
		Answer: pans,
	}
}

type Question struct {
	Id      int    `json:"id"`
	Type    string `json:"type" binding:"required"`
	Content any    `json:"content" binding:"required"`
}

func qTabEntity2Question(q *repository.Question, containsAnswer bool) *Question {
	var question any
	switch q.Type {
	case "single":
		question = qTabEntity2Single(q, containsAnswer)
	case "multiple":
		question = qTabEntity2Multiple(q, containsAnswer)
	case "fill":
		question = qTabEntity2Fill(q, containsAnswer)
	case "judge":
		question = qTabEntity2Judge(q, containsAnswer)
	default:
		panic("??????????????????????????????" + q.Type)
	}
	return &Question{
		Id:      q.Qid,
		Type:    q.Type,
		Content: question,
	}
}
func GetQuestion(operatorUid int, id int, containsAnswer bool) (*Question, error) {
	q, err := repository.GetQuestion(id)
	if err != nil {
		return nil, err
	}
	return qTabEntity2Question(q, containsAnswer), nil
}

func GetUserQuestions(uid int, pageId int, limit int) ([]Question, error) {
	rqs, err := repository.GetUserQuestions(uid, pageId, limit)
	if err != nil {
		return nil, err
	}
	var qs []Question
	for _, rq := range rqs {
		qs = append(qs, *qTabEntity2Question(&rq, true))
	}
	return qs, nil
}
