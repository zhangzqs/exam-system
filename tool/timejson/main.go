package main

import (
	"encoding/json"
	"log"
	"time"
)

func main() {
	now := time.Now()
	j, err := json.Marshal(now)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(j))
}
