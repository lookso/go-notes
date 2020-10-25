package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmdStr := fmt.Sprintf("ls -l %s | head -n %d", ".", 10)
	cmd := exec.Command("sh", "-c", cmdStr)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Errorf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %s\n%s\n", cmdStr, out)
	}
}
