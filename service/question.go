package service

import (
	"encoding/json"
	"github.com/zhangzqs/exam-system/repository"
	"strconv"
)

type SingleQuestion struct {
	Title   string
	Options []string
	Answer  int
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
	Title   string
	Options []string
	Answer  []int
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
	Title  string
	Answer []string
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
	Title  string
	Answer bool
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

func UpdateSingleQuestion(operatorUid int, id int, title string, options []string, answer int) int {
	return 0
}
func UpdateMultipleQuestion(operatorUid int, id int, title string, options []string, answer []int) int {
	return 0
}
func UpdateFillQuestion(operatorUid int, id int, title string, answer []string) int {
	return 0
}
func UpdateJudgeQuestion(operatorUid int, id int, title string, answer bool) int {
	return 0
}

func GetQuestion(operatorUid int, id int) {

}

func GetUserQuestions(uid int) {

}
