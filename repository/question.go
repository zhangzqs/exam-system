package repository

import (
	"github.com/zhangzqs/exam-system/global"
)

type Question struct {
	Uid     int
	Title   string
	Type    string
	Options string
	Answer  string
}

func AddQuestion(question *Question) (id int, err error) {
	db := global.GetDatabase()
	if err = db.QueryRow("INSERT INTO questions (created_by, title, type, option, answer) VALUES ($1, $2, $3, $4, $5) RETURNING qid",
		question.Uid,
		question.Title,
		question.Type,
		question.Options,
		question.Answer,
	).Scan(&id); err != nil {
		return
	}
	return
}
