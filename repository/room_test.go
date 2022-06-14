package repository

import (
	"log"
	"testing"
	"time"
)

func TestInsertRoom(t *testing.T) {
	start := time.Now()
	end := start.Add(time.Hour)
	room, err := InsertRoom(1, start, end)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(room)
}
