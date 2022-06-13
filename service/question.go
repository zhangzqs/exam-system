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
	Answer  int      `json:"answer"`
}

func AddSingleQuestion(operatorUid int, question *SingleQuestion) (int, error) {
	jsonStr, _ := json.Marshal(question.Options)
	answerIdStr := strconv.Itoa(question.Answer)
	qid, err := repository.AddQuestion(&repository.Question{
		Uid:     operatorUid,
		Title:   question.Title,
		Type:    "single",
		Options: string(jsonStr),
		Answer:  answerIdStr,
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
		Uid:     operatorUid,
		Title:   question.Title,
		Type:    "multiple",
		Options: string(jsonOptionsStr),
		Answer:  string(jsonAnswerStr),
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
		Uid:     operatorUid,
		Title:   question.Title,
		Type:    "fill",
		Options: "[]",
		Answer:  string(jsonAnswerStr),
	})

	if err != nil {
		return 0, err
	}
	return qid, nil
}

type JudgeQuestion struct {
	Title  string `json:"title"`
	Answer bool   `json:"answer"`
}

func AddJudgeQuestion(operatorUid int, question *JudgeQuestion) (int, error) {
	jsonAnswerStr, _ := json.Marshal(question.Answer)

	qid, err := repository.AddQuestion(&repository.Question{
		Uid:     operatorUid,
		Title:   question.Title,
		Type:    "judge",
		Options: "[]",
		Answer:  string(jsonAnswerStr),
	})

	if err != nil {
		return 0, err
	}
	return qid, nil
}

func DeleteQuestion(operatorUid int, id int) bool {
	return true
}

func UpdateSingleQuestion(operatorUid int, qid int, question *SingleQuestion) error {
	return nil
}
func UpdateMultipleQuestion(operatorUid int, qid int, question *MultipleQuestion) error {
	return nil

}
func UpdateFillQuestion(operatorUid int, qid int, question *FillQuestion) error {
	return nil

}
func UpdateJudgeQuestion(operatorUid int, qid int, question *JudgeQuestion) error {
	return nil

}

func GetQuestion(operatorUid int, id int) any {
	return nil
}

func GetUserQuestions(uid int) {

}
