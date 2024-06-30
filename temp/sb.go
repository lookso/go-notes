package main

import (
	"fmt"
	"strings"
	"unicode"
)

func splitText(text string) []string {
	var result []string

	charType := 0 // 1 for Chinese, 2 for English, 3 for punctuation
	temp := ""

	for _, char := range text {
		theType := 3
		if isChinese(char) {
			theType = 1
		} else if unicode.IsLetter(char) {
			theType = 2
		}
		if charType == 0 {
			charType = theType
			temp = string(char)
		} else if theType == 1 && charType == 1 {
			temp += string(char)
		} else if theType == 2 && charType == 2 {
			temp += string(char)
		} else if theType == 3 {
			temp += string(char)
		} else {
			result = append(result, temp)
			temp = string(char)
			charType = theType
		}
	}

	if temp != "" {
		result = append(result, temp)
	}

	return result
}

func isChinese(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

// func removeSpecialChars(s string) string {
// 	// 正则表达式匹配非字母和非数字字符
// 	reg := regexp.MustCompile(`[^a-zA-Z0-9\\u4e00-\\u9fa5]`)
// 	return reg.ReplaceAllString(s, "")
// }

func containsChineseOrEnglishOrDigit(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) || unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func main() {
	//txt := splitText("  the robot watched him sleeping. It knew everything about Mr. Moft.")
	txt := splitText(",,,, 马勒戈壁hello 213")
	for _, v := range txt {
		fmt.Printf("v:%+v,removeSpecialChars:%+v\n", v, containsChineseOrEnglishOrDigit(v))
		//fmt.Printf()
	}
	// arr := splitTextByLen("真牛逼啊,we are family,haha", 2)
	// fmt.Println(arr[0])
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }
}

func splitTextByLen(text string, length int) []string {
	var result []string
	var line string
	var charCount int

	words := strings.Fields(text)
	fmt.Println("words", words, len(words))

	for _, word := range words {
		wordLen := len([]rune(word))
		// 如果单词长度大于给定长度，直接添加到结果中，并开始新的行
		if wordLen > length {
			result = append(result, word)
			line = ""
			charCount = 0
		} else {
			// 如果当前行加上新单词会超过给定长度，添加当前行到结果中，并开始新的行
			if charCount+wordLen+1 > length { // 加1是因为单词之间有一个空格
				result = append(result, line)
				line = word
				charCount = wordLen
			} else {
				// 否则，将新单词添加到当前行
				if line != "" {
					line += " "
					charCount++
				}
				line += word
				charCount += wordLen
			}
		}
	}

	// 如果最后一行不为空，添加到结果中
	if line != "" {
		result = append(result, line)
	}

	return result
}
