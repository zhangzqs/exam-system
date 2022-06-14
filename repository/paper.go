package repository

import "github.com/zhangzqs/exam-system/global"

func AddPaper(uid int, title string) (pid int, err error) {
	db := global.GetDatabase()
	if err = db.QueryRow(
		"INSERT INTO papers (created_by, title) VALUES ($1, $2) RETURNING pid",
		uid,
		title,
	).Scan(&pid); err != nil {
		return
	}
	return
}

func PutPaperQuestion(qid int, pid int, score float64) error {
	db := global.GetDatabase()
	if _, err := db.Exec(
		"INSERT INTO paper_question (qid,pid,score) "+
			"VALUES ($1,$2,$3) "+
			"ON CONFLICT ON CONSTRAINT paper_question_pk "+
			"DO UPDATE SET score = excluded.score",
		qid,
		pid,
		score,
	); err != nil {
		return err
	}
	return nil
}

func GetPaperQuestion(qid int, pid int) (score float64, err error) {
	db := global.GetDatabase()
	if err = db.QueryRow(
		"SELECT score FROM paper_question WHERE qid=$1 AND pid=$2",
		qid,
		pid,
	).Scan(&score); err != nil {
		return
	}
	return
}
