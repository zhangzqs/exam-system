package repository

import "github.com/zhangzqs/exam-system/global"

func SubmitAnswer(uid int, rid int, qid int, answer string) error {
	db := global.GetDatabase()
	if _, err := db.Exec(
		"INSERT INTO user_answer(uid,rid,qid,user_answer) "+
			"VALUES ($1,$2,$3,$4) "+
			"ON CONFLICT ON CONSTRAINT user_answer_pk "+
			"DO UPDATE SET user_answer = excluded.user_answer",
		uid, rid, qid, answer,
	); err != nil {
		return err
	}
	return nil
}
