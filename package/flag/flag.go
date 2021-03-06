package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func getString() {
	var str = flag.String("name", "jack", "-name toms")
	flag.Parse()
	fmt.Println("str", *str)
}
func getInt() {
	var ip = flag.Int("flagname", 1234, "help message for flagname")

	flag.Parse()
	fmt.Println("ip has value ", *ip)
}

func main() {
	getInt()
	getString()

	var confPointer = flag.String("conf", "", "-conf ../conf/dev_config.toml")
	flag.Set("conf", "abc")
	fmt.Println(*confPointer)
	// 从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	fmt.Println(*confPointer)
	configFile, _ := filepath.Abs(*confPointer)
	if filepath.IsAbs(configFile) {
		fmt.Println(configFile)
	}
}
