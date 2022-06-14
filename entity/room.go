package entity

import "time"

type RoomEntity struct {
	RoomId    int       `json:"roomId"`
	PaperId   int       `json:"paperId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
