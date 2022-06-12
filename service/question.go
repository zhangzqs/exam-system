package service

func AddSingleQuestion(operatorUid int, title string, options []string, answer int) int {
	return 0
}
func AddMultipleQuestion(operatorUid int, title string, options []string, answer []int) int {
	return 0
}
func AddFillQuestion(operatorUid int, title string, answer []string) int {
	return 0
}
func AddJudgeQuestion(operatorUid int, title string, answer bool) int {
	return 0
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
