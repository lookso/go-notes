package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/peanut/apps/works/go/src/go-notes/project/download/stat.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var totalBandwidth float64
	var lineCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		bandwidthStr := strings.Split(parts[0], "-")[0]
		bandwidth, err := strconv.ParseFloat(bandwidthStr, 64)
		if err != nil {
			fmt.Println("Error parsing bandwidth:", err)
			return
		}
		totalBandwidth += bandwidth
		lineCount++
	}
	fmt.Println("lineCount", lineCount)

	averageBandwidth := totalBandwidth / float64(lineCount)
	fmt.Printf("Average bandwidth: %.2f-MB/sec\n", averageBandwidth)
}
