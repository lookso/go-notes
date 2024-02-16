package main

import (
	cmd2 "go-notes/libs/cli/cobra/tour/cmd"
	"log"
)

// go run main.go word -s=toms -m=1
func main() {
	err := cmd2.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
