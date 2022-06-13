package repository

import (
	"log"
	"testing"
)

func TestAddQuestion(t *testing.T) {
	qid, err := AddQuestion(&Question{
		Uid:     1,
		Title:   "title123",
		Type:    "single",
		Options: "options1",
		Answer:  "answer123",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("qid: ", qid)
}