package repository

import "github.com/zhangzqs/exam-system/global"

func UpdateEnterRoomTime(roomId int, studentId int) error {
	db := global.GetDatabase()
	if _, err := db.Exec(
		"UPDATE user_room "+
			"SET enter_at = CASE "+
			"    WHEN enter_at IS NULL THEN now() "+
			"    ELSE enter_at "+
			"END "+
			"WHERE rid = $1 "+
			"  AND uid = $2",
		roomId, studentId,
	); err != nil {
		return err
	}
	return nil
}
func UpdateSubmitTime(roomId int, studentId int) error {
	db := global.GetDatabase()
	if _, err := db.Exec(
		"UPDATE user_room "+
			"SET submit_at = now() "+
			"WHERE rid = $1 "+
			"  AND uid = $2",
		roomId, studentId,
	); err != nil {
		return err
	}
	return nil
}
func UpdateCommentAndScore(roomId int, studentId int, comment string, score float64) error {
	db := global.GetDatabase()
	if _, err := db.Exec(
		"UPDATE user_room "+
			"SET comment = $3, "+
			"	 score = $4"+
			"WHERE rid = $1 "+
			"  AND uid = $2",
		roomId, studentId, comment, score,
	); err != nil {
		return err
	}
	return nil
}

func GetCommentAndScore(roomId int, studentId int) (comment *string, score *float64, err error) {
	db := global.GetDatabase()
	if err := db.QueryRow(
		"SELECT comment,score FROM user_room WHERE rid=$1 AND uid=$2",
		roomId, studentId,
	).Scan(&comment, &score); err != nil {
		return nil, nil, err
	}
	return
}

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

func CountScore(uid int, rid int) (float64, error) {
	db := global.GetDatabase()

	var pid int
	if err := db.QueryRow("SELECT pid FROM rooms WHERE rid=$1", rid).Scan(&pid); err != nil {
		return 0, err
	}

	cur, err := db.Query(
		"SELECT qid,user_answer "+
			"FROM user_answer "+
			"WHERE uid=$1 "+
			"  AND rid=$2 ",
		uid, rid,
	)
	if err != nil {
		return 0, err
	}

	var sumScore float64

	for cur.Next() {
		var qid int
		var ans string

		if err := cur.Scan(&qid, &ans); err != nil {
			return 0, err
		}

		var realAns string
		if err := db.QueryRow("SELECT answer FROM questions WHERE qid=$1", qid).Scan(&realAns); err != nil {
			return 0, err
		}

		var score float64
		if err := db.QueryRow("SELECT score FROM paper_question WHERE qid=$1 AND pid=$2", qid, pid).Scan(&score); err != nil {
			return 0, err
		}

		if ans == realAns {
			sumScore += score

		}
	}
	return sumScore, nil
}
