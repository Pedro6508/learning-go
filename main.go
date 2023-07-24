package main

import (
	hospitalQueue "LearnGo/challenge4"
	"log"
	"os"
)

const (
	MAX_PROCS_NUM = 5
)

func main() {
	errA := os.Setenv("PORT", "8080")
	errB := os.Setenv("HTML_PATH", "challenge4/**/*.html")

	if errA == nil && errB == nil {
		hospitalQueue.GinServer()
	} else {
		if errA != nil {
			log.Fatal(errA)
		}

		log.Fatal(errB)
	}
}
