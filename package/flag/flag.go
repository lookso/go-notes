package main

import (
	"flag"
	"fmt"
)

func main() {

	//getInt()
	getString()
}

func getString() {
	var str = flag.String("name","jack","-name toms")
	flag.Parse()
	fmt.Println("str",*str)
}
func getInt() {
	var ip = flag.Int("flagname", 1234, "help message for flagname")

	flag.Parse()
	fmt.Println("ip has value ", *ip)
}
