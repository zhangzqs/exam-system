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

type PaperInfo struct {
	Pid       int
	CreatedBy int
	Title     string
}

func GetPaperInfo(pid int) (*PaperInfo, error) {
	db := global.GetDatabase()
	var pi PaperInfo
	if err := db.QueryRow(
		"SELECT pid,created_by,title "+
			"FROM papers "+
			"WHERE pid=$1",
		pid).Scan(&pi.Pid, &pi.CreatedBy, &pi.Title); err != nil {
		return nil, err
	}
	return &pi, nil
}

type PaperQuestion struct {
	Score float64
	Qid   int
}

func GetPaperQuestions(pid int) ([]PaperQuestion, error) {
	db := global.GetDatabase()
	cur, err := db.Query("SELECT qid,score FROM paper_question WHERE pid=$1", pid)
	if err != nil {
		return nil, err
	}
	var pqs []PaperQuestion
	for cur.Next() {
		var pq PaperQuestion
		err := cur.Scan(&pq.Qid, &pq.Score)
		if err != nil {
			return nil, err
		}
		pqs = append(pqs, pq)
	}
	return pqs, nil
}
