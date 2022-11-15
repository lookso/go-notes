package main

import (
	"go-notes/gopackages/cli/cobra/tour/cmd"
	"log"
)

// go run main.go word -s=toms -m=1
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
