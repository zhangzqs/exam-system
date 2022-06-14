package repository

import (
	"log"
	"testing"
)

func TestAddPaper(t *testing.T) {
	pid, err := AddPaper(1, "Paper123")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("pid: ", pid)
}

func TestPutPaperQuestion(t *testing.T) {

	if err := PutPaperQuestion(2, 1, 2); err != nil {
		log.Fatalln(err)
	}

	if err := PutPaperQuestion(3, 2, 2); err != nil {
		log.Fatalln(err)
	}

	if err := PutPaperQuestion(3, 2, 2); err != nil {
		log.Fatalln(err)
	}
}

func TestGetPaperQuestionScore(t *testing.T) {
	//score, err := GetPaperQuestion(3, 2)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(score)
}
