package repository

import (
	"log"
	"testing"
)

func TestAddQuestion(t *testing.T) {
	qid, err := AddQuestion(&Question{
		CreatedBy: 1,
		Title:     "title123",
		Type:      "single",
		Options:   "options1",
		Answer:    "answer123",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("qid: ", qid)
}

func TestGetQuestion(t *testing.T) {
	q, err := GetQuestion(2)
	if err != nil {
		log.Println(err)
	}
	log.Println(q)
}
