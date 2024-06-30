package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// downloadResource downloads the resource from the given URL and saves it with a specified fileName.
func downloadResource(URL, fileName string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	// 根据主键id+处理字段创建文件夹
	packagePath := "resources" // 文件夹路径
	if _, err := os.Stat(packagePath); os.IsNotExist(err) {
		// 如果文件夹不存在，则创建文件夹
		err := os.MkdirAll(packagePath, os.ModePerm) // 使用MkdirAll以递归创建所有必需的父目录
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	//使用os.Open打开1.html文件
	file, err := os.Open("html/4.html")
	if err != nil {
		fmt.Println("Error opening HTML file:", err)
		return
	}
	defer file.Close()

	// 使用io.ReadAll读取文件内容
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading HTML file:", err)
		return
	}
	htmlContent := string(fileContent) // 将文件内容转换为字符串

	// Regex to find all URLs
	urlRegex := regexp.MustCompile(`(http[s]?://[^"]+)`)

	// Find all URLs
	matches := urlRegex.FindAllStringSubmatch(htmlContent, -1)
	for _, match := range matches {
		url := match[0]
		// Simplified logic to generate a local file name based on the URL
		// You might want to create a more sophisticated naming mechanism
		splitURL := strings.Split(url, "/")
		fileName := packagePath + "/" + splitURL[len(splitURL)-1]
		if !strings.Contains(fileName, ".") {
			if strings.Contains(fileName, "ims?") {
				fileName = fileName + ".png"
			} else {
				fmt.Println("file no subfix Error fileName:%+v", fileName)
				return
			}
		}
		//fmt.Println("Downloading:", url, "to", fileName)
		if err := downloadResource(url, fileName); err != nil {
			fmt.Println("Error downloading file:", err)
			return
		}

		// 过滤-找关键字是否匹配
		// 找资源
		// 下载资源
		// 针对图片模糊
		// 针对图片暗水印
		// 全部 upload oss
		// 上传完成删本地文件夹

		htmlContent = strings.Replace(htmlContent, url, fileName, 1)
	}

	// Save the modified HTML content to a new file
	err = os.WriteFile("modified_1.html", []byte(htmlContent), 0644)
	if err != nil {
		fmt.Println("Error saving modified HTML file:", err)
		return
	}

	fmt.Println("Modified HTML document saved successfully.")
}
