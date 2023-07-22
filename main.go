package main

import (
	hospitalQueue "LearnGo/challenge4"
	"os"
)

func main() {
	os.Setenv("PORT", "8080")
	os.Setenv("HTML_PATH", "challenge4/**/*.html")
	hospitalQueue.GinServer()
}
