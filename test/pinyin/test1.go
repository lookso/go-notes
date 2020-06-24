package main

import (
	"fmt"
	"go-notes/test/pinyin/pinyin"
)

func main() {
	pinyin.LoadingPYFileName("./pinyin/pinyin.txt") //这里是字典文件路径程序启动调用一次，载入缓存
	//demo
	str1, err := pinyin.To_Py("你知道吗还", "", "") //默认造型： hanzipinyin
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str1)
	str2, err := pinyin.To_Py("汉字拼音", "", pinyin.Tone) //带声调：hànzìpīnyīn
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str2)
	str3, err := pinyin.To_Py("汉字拼音", "", pinyin.InitialsInCapitals) //首字母大写无声调：HanZiPinYin
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str3)
	str4, err := pinyin.To_Py("汉字拼音", "-", pinyin.InitialsInCapitals) //首字母大写无声调加-分割：Han-Zi-Pin-Yin
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str4)
}
