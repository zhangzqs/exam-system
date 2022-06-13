package repository

import (
	"github.com/zhangzqs/exam-system/global"
)

type Question struct {
	Qid       int
	CreatedBy int
	Title     string
	Type      string
	Options   string
	Answer    string
}

func AddQuestion(question *Question) (id int, err error) {
	db := global.GetDatabase()
	if err = db.QueryRow("INSERT INTO questions (created_by, title, type, option, answer) VALUES ($1, $2, $3, $4, $5) RETURNING qid",
		question.CreatedBy,
		question.Title,
		question.Type,
		question.Options,
		question.Answer,
	).Scan(&id); err != nil {
		return
	}
	return
}

func GetQuestion(id int) (*Question, error) {
	var q Question
	db := global.GetDatabase()
	if err := db.QueryRow("SELECT created_by, title, type, option, answer FROM questions WHERE qid=$1", id).Scan(
		&q.CreatedBy,
		&q.Title,
		&q.Type,
		&q.Options,
		&q.Answer); err != nil {
		return nil, err
	}
	return &q, nil
}

// GetUserQuestions 获取用户的所有题目
func GetUserQuestions(uid int, pageId int, limit int) ([]Question, error) {
	db := global.GetDatabase()
	cur, err := db.Query("SELECT * FROM questions WHERE created_by=$1 LIMIT=$2 OFFSET=$3", uid, pageId, limit)
	if err != nil {
		return nil, err
	}
	var qs []Question
	for cur.Next() {
		var q Question
		err := cur.Scan(&q.Qid, &q.CreatedBy, &q.Title, &q.Type, &q.Options, &q.Answer)
		if err != nil {
			return nil, err
		}
		qs = append(qs, q)
	}
	return qs, nil
}
