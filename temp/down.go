package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// downloadFileAndDetectType downloads the file and tries to detect its content type based on the file header.
func downloadFileAndDetectType(url, outputPath string) (detectedContentType string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	// Create a buffer to store the header for content type detection
	header := make([]byte, 512) // 512 bytes are enough for http.DetectContentType to work
	_, err = resp.Body.Read(header)
	if err != nil && err != io.EOF {
		return "", err
	}

	// Detect the content type
	detectedContentType = http.DetectContentType(header)

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Write the header bytes to the file
	_, err = outFile.Write(header)
	if err != nil {
		return "", err
	}

	// Copy the rest of the body
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return "", err
	}

	return detectedContentType, nil
}

func main() {
	url := "https://qz-server.oss-cn-beijing.aliyuncs.com/extproduce/202407/02/7caa399f-cc62-48fc-8472-f7c837fccba7"
	outputPath := "downloaded_file"

	contentType, err := downloadFileAndDetectType(url, outputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Detected Content-Type:", contentType)
}
