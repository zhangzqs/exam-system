package entity

import "time"

type Student struct {
	Uid      int       `json:"uid"`
	EnterAt  time.Time `json:"enterAt"`
	SubmitAt time.Time `json:"submitAt"`
	Comment  string    `json:"comment"`
	Score    float64   `json:"score"`
}
type RoomEntity struct {
	RoomId      int       `json:"roomId"`
	PaperId     int       `json:"paperId"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	StudentList []Student `json:"studentList"`
}
