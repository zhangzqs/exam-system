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

func UpdateComment(roomId int, studentId int, comment string) error {
	db := global.GetDatabase()
	if _, err := db.Exec(
		"UPDATE user_room "+
			"SET comment = $3 "+
			"WHERE rid = $1 "+
			"  AND uid = $2",
		roomId, studentId, comment,
	); err != nil {
		return err
	}
	return nil
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

func CountScore() {

}
