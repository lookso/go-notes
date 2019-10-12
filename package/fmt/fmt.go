package main

import "fmt"

/**
 * 输出绿色字符串
 * @params str	待输出的字符串
 */
func PrintGreen(str string) {
	printColor(str, 32)
}

/**
 * 输出红色字符串
 * @params str	待输出的字符串
 */
func PrintRed(str string) {
	printColor(str, 31)
}

/**
 * 输出待颜色字符串
 * @params str	待输出的字符串
 * @params color 待输出颜色
 */
func printColor(str string, color int32) {
	fmt.Printf("%c[0;0;%vm%s%c[0m\n", 0x1B, color, str, 0x1B)
}

func main() {
	PrintGreen("绿色植物向阳开？")
	PrintRed("血色浪漫再继续")

	fmt.Println("")

	// 前景 背景 颜色
	// ---------------------------------------
	// 30  40  黑色
	// 31  41  红色
	// 32  42  绿色
	// 33  43  黄色
	// 34  44  蓝色
	// 35  45  紫红色
	// 36  46  青蓝色
	// 37  47  白色
	//
	// 代码 意义
	// -------------------------
	//  0  终端默认设置
	//  1  高亮显示
	//  4  使用下划线
	//  5  闪烁
	//  7  反白显示
	//  8  不可见

	for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	}

}
