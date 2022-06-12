package service

import "testing"

func TestAddSingleQuestion(t *testing.T) {
	AddSingleQuestion(1, &SingleQuestion{
		"t123", []string{"o123"}, 0,
	})
}

func TestAddMultipleQuestion(t *testing.T) {
	AddMultipleQuestion(1, &MultipleQuestion{
		Title: "multiple123",
		Options: []string{
			"option1", "options2", "options3", "options4",
		},
		Answer: []int{0, 1, 2},
	})
}

func TestAddFillQuestion(t *testing.T) {
	AddFillQuestion(1, &FillQuestion{
		"fill123", []string{"f1"},
	})
}

func TestAddJudgeQuestion(t *testing.T) {
	AddJudgeQuestion(1, &JudgeQuestion{
		Title:  "judge123",
		Answer: false,
	})
}
