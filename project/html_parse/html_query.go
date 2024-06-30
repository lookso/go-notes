// func main() {
// 	// 使用os.Open打开1.html文件
// 	file, err := os.Open("html/2.html")
// 	if err != nil {
// 		fmt.Println("Error opening HTML file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// 使用io.ReadAll读取文件内容
// 	fileContent, err := io.ReadAll(file)
// 	if err != nil {
// 		fmt.Println("Error reading HTML file:", err)
// 		return
// 	}

// 	html := string(fileContent) // 将文件内容转换为字符串

// 	// Load the HTML document from a string
// 	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
// 	if err != nil {
// 		fmt.Println("Error loading HTML document:", err)
// 		return
// 	}
// 	packagePath := "resource" // 文件夹路径
// 	if _, err := os.Stat(packagePath); os.IsNotExist(err) {
// 		// 如果文件夹不存在，则创建文件夹
// 		err := os.MkdirAll(packagePath, os.ModePerm) // 使用MkdirAll以递归创建所有必需的父目录
// 		if err != nil {
// 			fmt.Println("Error creating directory:", err)
// 			return
// 		}
// 	}
// 	indexTag := 0
// 	// Find and download all images and audio files
// 	doc.Find("img, source").Each(func(index int, item *goquery.Selection) {
// 		findex := cast.ToString(index)
// 		var URL string
// 		var exists bool
// 		if item.Is("img") {
// 			URL, exists = item.Attr("src")
// 			findex = findex + ".png"
// 		} else if item.Is("source") {
// 			URL, exists = item.Attr("src")
// 			if strings.Contains(URL, ".mp3") {
// 				findex = findex + ".mp3"
// 			}
// 		}

// 		if exists && (strings.HasPrefix(URL, "http://") || strings.HasPrefix(URL, "https://")) {
// 			fileName := fmt.Sprintf("resource/%s", findex) // Generate a simple fileName. Consider parsing URL for a better name.
// 			fmt.Println("Downloading:", URL)
// 			if err := downloadResource(URL, fileName); err != nil {
// 				fmt.Println("Error downloading resource:", err)
// 			} else {
// 				item.SetAttr("src", fileName)
// 				fmt.Println("Downloaded", fileName)
// 			}
// 		}
// 		indexTag = index
// 	})

// 	// Find and download all CSS files
// 	doc.Find("link[rel='stylesheet']").Each(func(index int, item *goquery.Selection) {
// 		index = indexTag + index + 1
// 		cssHref, exists := item.Attr("href")
// 		if exists && (strings.HasPrefix(cssHref, "http://") || strings.HasPrefix(cssHref, "https://")) {
// 			fileName := fmt.Sprintf("resource/%d.css", index) // Simple fileName generation
// 			fmt.Println("Downloading:", cssHref)
// 			if err := downloadResource(cssHref, fileName); err != nil {
// 				fmt.Println("Error downloading CSS file:", err)
// 			} else {
// 				item.SetAttr("href", fileName)
// 				fmt.Println("Downloaded", fileName)
// 			}
// 		}
// 	})
// 	// Save the modified HTML document
// 	modifiedHTML, err := doc.Html()
// 	if err != nil {
// 		fmt.Println("Error generating modified HTML:", err)
// 		return
// 	}

// 	err = os.WriteFile("modified_3.html", []byte(modifiedHTML), 0644)
// 	if err != nil {
// 		fmt.Println("Error saving modified HTML file:", err)
// 		return
// 	}
// }