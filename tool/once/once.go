package main

import (
	"log"
	"sync"
)

func main() {
	var o sync.Once
	for i := 0; i < 12; i++ {
		o.Do(func() {
			log.Println("执行了")
		})
	}
}
