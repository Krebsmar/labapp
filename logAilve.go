package main

import (
	"log"
	"time"
)

func logAlive() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		log.Println("I'm alive!")
		log.Printf("still running at: %v", time.Now().Format("2006-01-02 15:04:05"))
		<-ticker.C
	}
}
